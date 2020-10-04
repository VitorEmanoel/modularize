package events

import (
	"log"
	"testing"
	"time"
)

var testFunc = func (teste string) {
	log.Println("From event method >>", teste)
}

var (
	SimpleEvent = "events/SimpleEvents"
)

func TestSimpleEvent(t *testing.T) {
	manager := NewEventManager()
	err := manager.RegisterEvent(SimpleEvent, testFunc)
	if err != nil {
		t.Error("Error in register event", err.Error())
		return
	}
	time.Sleep(1 * time.Second)
	err = manager.CallEvent(SimpleEvent, "vitor")
	if err != nil {
		t.Error("Error in call event after 1 second")
		return
	}
	t.Log("Successful call simple event")
}