package modularize

import (
	"context"
	"log"
	"modularize/events"
	"time"
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
	EventManager 	events.EventManager
	PreModules		[]Module
	PreExtensions 	[]Extension
	Modules			[]moduleMapping
	Extensions		[]extensionMapping
	AppData			map[string]interface{}
	Resources		*Resources
	Options			AppOptions
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
	err := a.EventManager.CallEvent(OnDisableEvent)
	if err != nil {
		panic(err)
	}
}

func (a *appContext) Start(modules ...Module) {
	a.PreModules = append(a.PreModules, modules...)
	a.start()
}

// Private methods

func (a *appContext) start() {
	for _, extension := range a.PreExtensions {
		a.addExtension(extension)
	}
	for _, module := range a.PreModules {
		a.AddModules(module)
	}
	err := a.EventManager.CallEvent(OnEnableEvent)
	if err != nil {
		panic(err)
	}
}

func (a *appContext) load() {

}

func (a *appContext) initializeExtension(context context.Context, extension Extension, manager ExtensionManager) {
	extension(manager)
	select {
	case <-context.Done():
		log.Fatalln("Timeout on initialize extension")
	default:
		break
	}
}

func (a *appContext) initializeModule(context context.Context, module Module, manager ModuleManager) {
	module(manager)
	select {
	case <-context.Done():
		log.Fatalln("Timeout on initialize module")
	default:
		break
	}
}

func (a *appContext) addModule(module Module) {
	manager := NewModuleManager(a.EventManager, a.Resources)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	a.initializeModule(ctx, module, manager)
	a.Modules = append(a.Modules, moduleMapping{
		Module:  module,
		Manager: manager,
	})
}

func (a *appContext) addExtension(extension Extension) {
	manager := NewExtensionManager(a.Resources)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	a.initializeExtension(ctx, extension, manager)
	a.Extensions = append(a.Extensions, extensionMapping{
		Extension: extension,
		Manager:   manager,
	})

}

func Default() App {
	return &appContext{
		Options: AppOptions{
			ModulePath: "./modules",
			ExtensionPath: "./extensions",
		},
		EventManager: events.NewEventManager(),
		AppData: make(map[string]interface{}),
		Resources: NewResources(),
	}
}

func New(options AppOptions) App {
	return &appContext{
		Options: options,
		EventManager: events.NewEventManager(),
		AppData: make(map[string]interface{}),
		Resources: NewResources(),
	}
}

