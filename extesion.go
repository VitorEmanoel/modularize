package modularize

type Extension func (ctx ExtensionManager)

type ExtensionInfo struct {
	Name		string
}

type ExtensionManager interface {
	SetResource(name string, data interface{})
	SetInfo(info ExtensionInfo)
	GetInfo() ExtensionInfo
}

type ExtensionContext struct {
	Resources 	*Resources
	Info		ExtensionInfo
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

func NewExtensionManager(resources *Resources) ExtensionManager {
	return &ExtensionContext{
		Resources: resources,
	}
}
