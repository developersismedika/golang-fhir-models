package fhir

import (
	"encoding/json"
	"reflect"
)

type BaseResource struct {
	Type         reflect.Type
	ResourceType ResourceType
	Resource     interface{}
}

type resourceDefinition struct {
	ResourceType ResourceType `json:"resourceType"`
}

func (r BaseResource) MarshalJSON() ([]byte, error) {
	def := reflect.StructOf([]reflect.StructField{
		reflect.StructField{
			Name:      r.Type.Name(),
			Type:      r.Type,
			Anonymous: true,
		},
		reflect.StructField{
			Name: "ResourceType",
			Type: reflect.TypeOf(ResourceTypeAccount),
			Tag:  "json:\"resourceType\"",
		},
	})

	payload := reflect.New(def)
	payload.Elem().FieldByName(r.Type.Name()).Set(reflect.ValueOf(r.Resource).Elem())
	payload.Elem().FieldByName("ResourceType").Set(reflect.ValueOf(r.ResourceType))

	// result := reflect.ValueOf(payload).MethodByName("MarshalJSON").Call([]reflect.Value{})
	result := reflect.ValueOf(json.Marshal).Call([]reflect.Value{payload})

	err, ok := result[1].Interface().(error)
	if ok && err != nil {
		return nil, err
	}

	return result[0].Bytes(), err
}

func (r *BaseResource) UnmarshalJSON(data []byte) error {
	var def resourceDefinition
	err := json.Unmarshal(data, &def)
	if err != nil {
		return err
	}

	resource := reflect.New(def.ResourceType.Type())

	result := reflect.ValueOf(json.Unmarshal).Call([]reflect.Value{reflect.ValueOf(data), resource})
	err, ok := result[0].Interface().(error)
	if ok && err != nil {
		return err
	}

	r.Type = def.ResourceType.Type()
	r.ResourceType = def.ResourceType
	r.Resource = resource.Interface()

	return nil
}
