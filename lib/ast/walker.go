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
		valueType := val.Type().Field(i)
		valueKind := valueField.Kind()
		if valueKind == reflect.String || valueKind == reflect.Bool || valueType.Name == "Node" {
			continue
		}
		if valueKind == reflect.Slice {
			for i := 0; i < valueField.Len(); i++ {
				if valueField.Index(i).Kind() == reflect.String {
					continue
				}
				node := valueField.Index(i).Interface().(INode)
				if node == nil {
					continue
				}
				node.SetParent(ast)
				name := valueField.Index(i).Type().String()[5:]
				this.callDelegate("Before", valueField.Index(i))
				this.callDelegate("Each", valueField.Index(i))
				this.callDelegate(name, valueField.Index(i))
				valueField.Index(i).Set(reflect.ValueOf(this.Walk(node)))
				this.callDelegate("After", valueField.Index(i))
			}
			continue
		}
		if valueField.IsNil() {
			continue
		}
		name := valueField.Type().String()[5:]
		node := valueField.Interface().(INode)
		node.SetParent(ast)
		this.callDelegate("Before", valueField)
		this.callDelegate("Each", valueField)
		this.callDelegate(name, valueField)
		val.Field(i).Set(reflect.ValueOf(this.Walk(node)))
		this.callDelegate("After", valueField)
	}
	return ast
}
