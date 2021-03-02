package modularize

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"

	"modularize/events"
	"modularize/plugins"
)

type AppOptions struct {
	ModulePath		string		`json:"modulePath"`
	ExtensionPath	string		`json:"extensionPath"`
}

type App interface {
	AddModules(modules ...Module)
	AddExtensions(extensions ...Extension)
	Inject(name string, data interface{})
	Start(modules ...Module)
	Stop()
}

type moduleMapping struct {
	Module		Module
	Manager		ModuleManager
}

type extensionMapping struct {
	Extension	Extension
	Manager		ExtensionManager
}

type appContext struct {
	ModulesEventManager     events.EventManager
	ExtensionsEventManager  events.EventManager
	PreModules              []Module
	PreExtensions           []Extension
	Modules                 []moduleMapping
	Extensions              []extensionMapping
	extensionWait           sync.WaitGroup
	moduleWait              sync.WaitGroup
	AppData                 map[string]interface{}
	Resources               *Resources
	Options                 AppOptions
	PluginLoader            plugins.PluginLoader
}

// Public methods

func (a *appContext) AddExtensions(extensions ...Extension) {
	a.PreExtensions = append(a.PreExtensions, extensions...)
}

func (a *appContext) AddModules(modules ...Module) {
	a.PreModules = append(a.PreModules, modules...)
}

func (a *appContext) Inject(name string, data interface{}) {
	a.AppData[name] = data
}

func (a *appContext) Stop() {
	logrus.Info("Stopping modules")
	err := a.ModulesEventManager.CallEvent(OnDisableEvent)
	if err != nil {
		logrus.Error("Error in stopping module. Error: ", err.Error())
		panic(err)
	}
	logrus.Info("All modules stopped")
}

func (a *appContext) setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	if err := a.ModulesEventManager.CallEvent(OnDisableEvent); err != nil {
		panic(err)
	}
	if err := a.ExtensionsEventManager.CallEvent(OnDisableEvent); err != nil {
		panic(err)
	}
}

func (a *appContext) Start(modules ...Module) {
	a.PreModules = append(a.PreModules, modules...)
	a.load()
	a.start()
	a.setupCloseHandler()
}

// Private methods

func (a *appContext) start() {
	logrus.Info("Starting extensions...")
	for _, extension := range a.PreExtensions {
		a.extensionWait.Add(1)
		go a.addExtension(extension)
	}
	a.extensionWait.Wait()
	logrus.Info("All extensions initialized")
	if err := a.ExtensionsEventManager.CallEvent(OnEnableEvent); err != nil && len(a.PreExtensions) > 0 {
		panic(err)
	}
	logrus.Info("Starting modules...")
	for _, module := range a.PreModules {
		a.moduleWait.Add(1)
		go a.addModule(module)
	}
	a.moduleWait.Wait()
	logrus.Info("All modules initialized")
	if err := a.ModulesEventManager.CallEvent(OnEnableEvent); err != nil && len(a.PreModules) > 0 {
		panic(err)
	}
}

func (a *appContext) loadModule(plugin plugins.Plugin) {
	symbol, err := plugin.FindSymbol("Module")
	if err != nil {
		panic(err)
	}
	module, ok := symbol.(Module)
	if !ok {
		return
	}
	a.PreModules = append(a.PreModules, module)
}

func (a *appContext) loadExtension(plugin plugins.Plugin) {
	symbol, err := plugin.FindSymbol("Extension")
	if err != nil {
		panic(err)
	}
	extension, ok := symbol.(Extension)
	if !ok {
		return
	}
	a.PreExtensions = append(a.PreExtensions, extension)
}

func (a *appContext) load() {
	extensions := a.PluginLoader.LoadFolder(a.Options.ExtensionPath)
	for _, extension := range extensions {
		a.loadExtension(extension)
	}
	modules := a.PluginLoader.LoadFolder(a.Options.ModulePath)
	for _, module := range modules {
		a.loadModule(module)
	}
}

func (a *appContext) initializeExtension(extension Extension, manager ExtensionManager) {
	extension(manager)
	a.moduleWait.Done()
	logrus.Info(manager.GetInfo().Name, " extension initialized successful.")
}

func (a *appContext) initializeModule(module Module, manager ModuleManager) {
	module(manager)
	a.moduleWait.Done()
	logrus.Info(manager.GetInfo().Name, " module initialized successful.")
}



func (a *appContext) addModule (module Module) {
	manager := NewModuleManager(a.ModulesEventManager, a.Resources)
	a.initializeModule(module, manager)
	a.Modules = append(a.Modules, moduleMapping{
		Module:  module,
		Manager: manager,
	})
}

func (a *appContext) addExtension(extension Extension) {
	manager := NewExtensionManager(a.Resources, a.ExtensionsEventManager)
	a.initializeExtension(extension, manager)
	a.Extensions = append(a.Extensions, extensionMapping{
		Extension: extension,
		Manager:   manager,
	})

}

func context() appContext {
	return appContext{
		ModulesEventManager:    events.NewEventManager(),
		ExtensionsEventManager: events.NewEventManager(),
		AppData:                make(map[string]interface{}),
		Resources:              NewResources(),
		PluginLoader:           plugins.NewPluginLoader(),
	}
}

func Default() App {
	var app = context()
	app.Options = AppOptions{
		ModulePath: "./modules",
		ExtensionPath: "./extensions",
	}
	return &app
}

func New(options AppOptions) App {
	var app = context()
	app.Options = options
	return &app
}

