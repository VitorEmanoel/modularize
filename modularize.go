package modularize

import (
	"context"
	"log"
	"modularize/events"
	"time"
)

type App interface {
	Start(modules ...Module)
	AddModule(module Module)
	AddModules(modules ...Module)
	Inject(name string, data interface{})
	Stop()
}

type moduleMapping struct {
	Module		Module
	Manager		ModuleManager
}

type appContext struct {
	EventManager events.EventManager
	Modules		[]moduleMapping
	AppData		map[string]interface{}
	Resources	*Resources
}


func (a *appContext) Stop() {
	err := a.EventManager.CallEvent(OnDisableEvent)
	if err != nil {
		panic(err)
	}
}

func (a *appContext) AddModule(module Module) {
	manager := NewModuleManager(a.EventManager, a.Resources)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	a.initialize(ctx, module, manager)
	a.Modules = append(a.Modules, moduleMapping{
		Module:  module,
		Manager: manager,
	})
}

func (a *appContext) AddModules(modules ...Module) {
	for _, module := range modules {
		a.AddModule(module)
	}
}

func (a *appContext) start() {
	err := a.EventManager.CallEvent(OnEnableEvent)
	if err != nil {
		panic(err)
	}
}

func (a *appContext) initialize(context context.Context, module Module, manager ModuleManager) {
	module(manager)
	select {
		case <-context.Done():
			log.Fatalln("Timeout on initialize module")
		default:
			break
	}
}

func (a *appContext) Start(modules ...Module) {
	a.AddModules(modules...)
	a.start()
}

func (a *appContext) Inject(name string, data interface{}) {
	a.AppData[name] = data
}

func Default() App {
	return &appContext{
		EventManager: events.NewEventManager(),
		AppData: make(map[string]interface{}),
		Resources: NewResources(),
	}
}
