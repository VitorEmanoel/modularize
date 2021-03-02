package modularize

import (
	"log"
	"testing"
	"time"
)

var TestModule = func (manager ModuleManager) {
	manager.OnEnable(func() {
		log.Println("Started Test Module")
	})

	manager.OnDisable(func () {
		log.Println("Disabled Test Module")
	})

	time.Sleep(10 * time.Second)
	manager.SetInfo(ModuleInfo{
		Name:    "TestModule",
		Version: "v1.0.0",
	})
}

func TestModulerizeStart(t *testing.T) {
	app := Default()
	app.Start(TestModule)
}
