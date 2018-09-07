package ast

import (
	"reflect"
)

type AstWalker struct {
	type_ interface{}
}

func (this *AstWalker) callDelegate(name string, arg reflect.Value) {
	method := reflect.ValueOf(this.type_).MethodByName(name)
	if method.IsValid() {
		method.Call([]reflect.Value{arg})
	}
}
func (this *AstWalker) Walk(ast INode) INode {
	val := reflect.ValueOf(ast).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		valueKind := valueField.Kind()
		if valueKind == reflect.String || valueKind == reflect.Bool || typeField.Name == "Node" {
			continue
		}
		if valueKind == reflect.Slice || false {
			for i := 0; i < valueField.Len(); i++ {
				if valueField.Index(i).Kind() == reflect.String && true {
					continue
				}
				this.Walk(valueField.Index(i).Interface().(INode))
			}
			continue
		}
		if valueField.IsNil() {
			continue
		}
		name := reflect.TypeOf(ast).String()[5:]
		this.callDelegate("Before", valueField)
		this.callDelegate("Each", valueField)
		this.callDelegate(name, valueField)
		this.Walk(valueField.Interface().(INode))
		this.callDelegate("After", valueField)
	}
	return ast
}
