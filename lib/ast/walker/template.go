package walker

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/champii/og/lib/common"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
)

type Template struct {
	Name         string
	Pack         string
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
func NewTemplate(name, pack string, types []string, node common.INode) *Template {
	return &Template{
		Name:         name,
		Pack:         pack,
		Types:        types,
		UsedFor:      [][]string{},
		Node:         node,
		GeneratedFor: make(map[string][]string),
	}
}

type Templates struct {
	Names     []string
	Packages  []string
	Templates []*Template
}

func (this *Templates) Add(name, pack string, template *Template) {
	this.Names = append(this.Names, name)
	this.Packages = append(this.Packages, pack)
	this.Templates = append(this.Templates, template)
}
func (this *Templates) Get(name, pack string) *Template {
	for i, n := range this.Names {
		if n == name && this.Packages[i] == pack {
			return this.Templates[i]
		}
	}
	return nil
}
func (this Templates) ResetUsedFor() {
	for _, template := range this.Templates {
		template.UsedFor = [][]string{}
	}
}
func (this *Templates) Decode(content []byte) {
	arr := []*TemplateSerie{}
	buf := bytes.NewBuffer(content)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&arr); err != nil {
		panic("ERROR DECODE" + err.Error())
	}
	for _, tmpl := range arr {
		this.Add(tmpl.Name, tmpl.Template.Pack, tmpl.Template)
	}
}
func (this *Templates) Encode(arr []*Template) []byte {
	res := []*TemplateSerie{}
	for _, template := range arr {
		res = append(res, &TemplateSerie{
			template.Name,
			template,
		})
	}
	var (
		buf bytes.Buffer
	)
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&res); err != nil {
		panic("ERROR ENCODE" + err.Error())
	}
	return buf.Bytes()
}
func (this *Templates) byPackage() map[string][]*Template {
	res := make(map[string][]*Template)
	for _, template := range this.Templates {
		res[template.Pack] = append(res[template.Pack], template)
	}
	return res
}
func (this Templates) Store() {
	for pack, arr := range this.byPackage() {
		templateDir := path.Join(pack, ".og")
		_, err := os.Stat(templateDir)
		if err != nil {
			err = os.Mkdir(templateDir, 0755)
			if err != nil {
				fmt.Println("Cannot create template directory: " + templateDir)
				return
			}
		}
		blob := this.Encode(arr)
		ioutil.WriteFile(path.Join(templateDir, "template"), blob, 0644)
	}
}
func NewTemplates() *Templates {
	return &Templates{}
}

type TemplateSerie struct {
	Name     string
	Template *Template
}
