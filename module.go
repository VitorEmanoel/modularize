package modularize

import "modularize/events"

type ModuleInfo struct {
	Name	string
	Version	string
}

type Modules []Module

type ModuleManager interface {
	OnEnable(executor OnEnableExecutor)
	OnDisable(executor OnDisableExecutor)
	SetInfo(info ModuleInfo)
	GetInfo() ModuleInfo
	Inject(data interface{})
}

type Module func (ctx ModuleManager)

type ModuleContext struct {
	EventManager 	events.EventManager
	Info			ModuleInfo
	Resources		*Resources
}

func (m *ModuleContext) Inject(data interface{}) {
	m.Resources.Inject(data)
}

func (m *ModuleContext) SetInfo(info ModuleInfo) {
	m.Info = info
}

func (m *ModuleContext) GetInfo() ModuleInfo {
	return m.Info
}

func (m *ModuleContext) OnEnable(executor OnEnableExecutor) {
	err := m.EventManager.RegisterEvent(OnEnableEvent, executor)
	if err != nil {
		panic(err)
	}
}

func (m *ModuleContext) OnDisable(executor OnDisableExecutor) {
	err := m.EventManager.RegisterEvent(OnDisableEvent, executor)
	if err != nil {
		panic(err)
	}
}

func NewModuleManager(eventManager events.EventManager, resources *Resources) ModuleManager {
	return &ModuleContext{
		EventManager: eventManager,
		Resources: resources,
	}
}