package modularize

import (
	"modularize/events"
)

type Extension func (ctx ExtensionManager)

type ExtensionInfo struct {
	Name		string
}

type ExtensionManager interface {
	OnEnable(executor OnEnableExecutor)
	OnDisable(executor OnDisableExecutor)
	SetResource(name string, data interface{})
	SetInfo(info ExtensionInfo)
	GetInfo() ExtensionInfo
}

type ExtensionContext struct {
	EventManager    events.EventManager
	Resources       *Resources
	Info		    ExtensionInfo
}

func (e *ExtensionContext) OnEnable(executor OnEnableExecutor) {
	err := e.EventManager.RegisterEvent(OnEnableEvent, executor)
	if err != nil {
		panic(err)
	}
}

func (e *ExtensionContext) OnDisable(executor OnDisableExecutor) {
	panic("implement me")
}

func (e *ExtensionContext) GetInfo() ExtensionInfo {
	return e.Info
}

func (e *ExtensionContext) SetResource(name string, data interface{}) {
	e.Resources.SetResource(name, data)
}

func (e *ExtensionContext) SetInfo(info ExtensionInfo) {
	e.Info = info
}

func NewExtensionManager(resources *Resources, eventManager events.EventManager) ExtensionManager {
	return &ExtensionContext{
		Resources: resources,
		EventManager: eventManager,
	}
}
