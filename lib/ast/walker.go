package ast

import (
	"reflect"
)

type AstWalker struct {
	type_ interface{}
}

func (this *AstWalker) callDelegate(name string, arg reflect.Value) ([]reflect.Value, bool) {
	method := reflect.ValueOf(this.type_).MethodByName(name)
	if method.IsValid() {
		res := method.Call([]reflect.Value{arg})
		if len(res) == 0 {
			return []reflect.Value{reflect.Zero(arg.Type())}, false
		} else {
			return res, true
		}
	}
	return []reflect.Value{reflect.Zero(arg.Type())}, false
}
func (this *AstWalker) Trigger(arg reflect.Value, parentField reflect.Value, parentNode INode) (reflect.Value, bool) {
	node := arg.Interface().(INode)
	if node == nil {
		return reflect.Zero(arg.Type()), false
	}
	node.SetParent(parentNode)
	name := arg.Type().String()[5:]
	this.callDelegate("Before", arg)
	this.callDelegate("Before"+name, arg)
	res, ok := this.callDelegate("Each", arg)
	if ok {
		node = res[0].Interface().(INode)
		parentField.Set(reflect.ValueOf(node))
	}
	this.callDelegate(name, arg)
	parentField.Set(reflect.ValueOf(this.Walk(node)))
	res, ok = this.callDelegate("After", arg)
	if ok {
		node = res[0].Interface().(INode)
		parentField.Set(reflect.ValueOf(node))
	}
	this.callDelegate("After"+name, arg)
	return reflect.ValueOf(node), true
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
			for j := 0; j < valueField.Len(); j++ {
				if valueField.Index(j).Kind() == reflect.String {
					continue
				}
				if !valueField.Index(j).IsValid() {
					continue
				}
				res, ok := this.Trigger(valueField.Index(j), valueField.Index(j), ast)
				if ok {
					valueField.Index(j).Set(res)
				}
			}
			continue
		}
		if valueField.IsNil() {
			continue
		}
		res, ok := this.Trigger(valueField, val.Field(i), ast)
		if ok {
			val.Field(i).Set(res)
		}
	}
	return ast
}
