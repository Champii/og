package walker

import (
	"github.com/champii/og/lib/common"
	"reflect"
	"strings"
)

type Template struct {
	Name         string
	Types        []string
	UsedFor      [][]string
	GeneratedFor map[string][]string
	Node         common.INode
}

func (this Template) contains(arr []string, str string) bool {
	for _, item := range arr {
		if str == item {
			return true
		}
	}
	return false
}
func (this *Template) IsGeneratedFor(types []string, packageName string) bool {
	serie := strings.Join(types, ",")
	packages, _ := this.GeneratedFor[serie]
	return this.contains(packages, packageName)
}
func (this *Template) AddGeneratedFor(types []string, packageName string) {
	serie := strings.Join(types, ",")
	packages, _ := this.GeneratedFor[serie]
	if !this.contains(packages, packageName) {
		this.GeneratedFor[serie] = append(this.GeneratedFor[serie], packageName)
	}
}
func (this *Template) AddUsedFor(types []string) {
	for _, usedFor := range this.UsedFor {
		if reflect.DeepEqual(types, usedFor) {
			return
		}
	}
	this.UsedFor = append(this.UsedFor, types)
}
func NewTemplate(name string, types []string, node common.INode) *Template {
	return &Template{
		Name:         name,
		Types:        types,
		UsedFor:      [][]string{},
		Node:         node,
		GeneratedFor: make(map[string][]string),
	}
}

type Templates struct {
	names     []string
	templates []*Template
}

func (this *Templates) Add(name string, template *Template) {
	this.names = append(this.names, name)
	this.templates = append(this.templates, template)
}
func (this *Templates) Get(name string) *Template {
	for i, n := range this.names {
		if n == name {
			return this.templates[i]
		}
	}
	return nil
}
func (this Templates) ResetUsedFor() {
	for _, template := range this.templates {
		template.UsedFor = [][]string{}
	}
}
func NewTemplates() *Templates {
	return &Templates{}
}
