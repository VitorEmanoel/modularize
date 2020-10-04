package events

import "reflect"

type eventMapping struct {
	FuncType		reflect.Type
	FuncValue 		reflect.Value
}

type EventManager interface {
	RegisterEvent(event string, method interface{}) error
	CallEvent(event string, args ...interface{}) error
}

type eventManagerContext struct {
	Events		map[string] []eventMapping
}

func (e *eventManagerContext) RegisterEvent(event string, method interface{}) error {
	funcType := reflect.TypeOf(method)
	if funcType.Kind() != reflect.Func {
		return MethodNotFunc
	}
	e.Events[event] = append(e.Events[event], eventMapping{
		FuncType:  funcType,
		FuncValue: reflect.ValueOf(method),
	})
	return nil
}

func (e *eventManagerContext) CallEvent(event string, args ...interface{}) error{
	events, exists := e.Events[event]
	if !exists {
		return EventNotFound
	}
	var methodValues []reflect.Value
	for _, arg := range args {
		methodValues = append(methodValues, reflect.ValueOf(arg))
	}
	for _, method := range events {
		if method.FuncType.NumIn() > len(methodValues) {
			return MissingMethodArgs
		}
		method.FuncValue.Call(methodValues)
	}
	return nil
}

func NewEventManager() EventManager {
	return &eventManagerContext{Events: make(map[string][]eventMapping)}
}