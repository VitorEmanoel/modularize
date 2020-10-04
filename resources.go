package modularize

import "reflect"

type Resources struct {
	Data		map[string]interface{}
}

func (r *Resources) SetResource(name string, data interface{}) {
	if r.Data == nil {
		r.Data = make(map[string]interface{})
	}
	r.Data[name] = data
}

func (r *Resources) Inject(data interface{}) {
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	if dataType.Kind() != reflect.Ptr {
		panic(NoPointErr)
	}
	elemType := dataType.Elem()
	elemValue := dataValue.Elem()
	if elemType.Kind() != reflect.Struct {
		panic(InvalidTypeErr)
	}
	for _, item := range r.Data {
		itemType := reflect.TypeOf(item)
		if itemType.AssignableTo(elemType) {
			elemValue.Set(reflect.ValueOf(item))
			return
		}
	}
	for i := 0; i < elemType.NumField(); i++{
		fieldType := elemType.Field(i)
		if injectName := fieldType.Tag.Get("inject"); injectName != "" && r.Data[injectName] != nil {
			resourceValue := r.Data[injectName]
			resourceType := reflect.TypeOf(resourceValue)
			if resourceType.AssignableTo(fieldType.Type) {
				dataValue.Elem().Field(i).Set(reflect.ValueOf(resourceValue))
			}
		}
	}
}

func NewResources() *Resources {
	return &Resources{Data: make(map[string]interface{})}
}
