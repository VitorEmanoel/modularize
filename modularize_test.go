package modularize

import (
	"testing"
)

var TestModule = func (manager ModuleManager) {
	manager.OnEnable(func() {
	})

	manager.OnDisable(func () {
	})

	manager.SetInfo(ModuleInfo{
		Name:    "TestModule",
		Version: "v1.0.0",
	})
}

func TestModulerizeStart(t *testing.T) {
	app := Default()
	app.Start(TestModule)
}
