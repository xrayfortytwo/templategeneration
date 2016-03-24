package genjson

/*
   Generates the json representation for the grammar to fill the template
   after marshalling
*/

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

type kind int

const (
	Plain kind = iota
	Object
	Array
)

type object struct {
	id     string
	parent string
	VType  kind
}

type Encoding struct {
	data  map[string]interface{}
	path  map[string][]string
	desti string
}

func (e Encoding) Write(p []byte) (int, error) {
	f, err := os.OpenFile(e.desti, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	check := func(err error, msg string) {
		if err != nil {
			log.Fatal("ERROR: ", e.desti, err)
		}
	}
	check(err, "file creation")
	_, err = f.Write(p)
	check(err, "write to file")
	log.Println("json written to file datajson.json")
	f.Close()
	return 1, nil
}

func GenJson(desti *string) *Encoding {
	return &Encoding{make(map[string]interface{}), make(map[string][]string), *desti}
}

func (e *Encoding) Add(id, parent string, VType kind) {
	e.savePath(e.addObject(&object{id, parent, VType}))
}

func (e *Encoding) Decode() {
	json.NewEncoder(e).Encode(e.data)
}

func (e *Encoding) addObject(o *object) *object {
	if _, ok := e.path[o.parent]; ok && o.parent != "" {
		s := e.searchPath(o.parent)
		if o.VType == Object {
			d := s[o.parent].(map[string]interface{})
			d[o.id] = make(map[string]interface{})
		} else if o.VType == Array {
			d := s[o.parent].(map[string][]interface{})
			d[o.id] = make([]interface{}, 0)
		} else if o.VType == Plain {
			d := s[o.parent].(map[string]interface{})
			d[o.id] = ""
		}
		return o
	}

	if e.checkForAlias(o) {
		return o
	}

	if _, ok := e.data[o.parent]; ok {
		switch o.VType {
		case Object:
			s := e.data[o.parent].(map[string]interface{})
			if len(s) < 1 {
				s = make(map[string]interface{})
			}
			s[o.id] = ""
			e.data[o.parent] = s
		case Array:
			a := e.data[o.parent].(map[string][]interface{})
			if len(a) < 1 {
				a = make(map[string][]interface{})
			}
			a[o.id] = make([]interface{}, 0)
			e.data[o.parent] = a
		}
	} else {
		switch o.VType {
		case Object:
			e.data[o.id] = make(map[string]interface{}, 0)
		case Array:
			e.data[o.id] = make([]interface{}, 0)
		case Plain:
			e.data[o.id] = ""
		}
	}
	return o
}

func (e *Encoding) searchPath(id string) map[string]interface{} {
	var z = e.data
	for _, s := range e.path[id] {
		if i, ok := e.furtherPath(s, z); ok {
			z = i
		}
	}
	return z
}

func (e *Encoding) furtherPath(
	id string,
	m map[string]interface{}) (map[string]interface{}, bool) {

	if i, ok := m[id]; ok {
		switch i.(type) {
		case (map[string]interface{}):
			return i.(map[string]interface{}), true
		case string:
			return m, false
		default:
			panic(reflect.TypeOf(i))
		}
	}
	return nil, false
}

func (e *Encoding) savePath(o *object) {
	for _, k := range e.path[o.parent] {
		e.path[o.id] = append(e.path[o.id], k)
	}
	e.path[o.id] = append(e.path[o.id], o.parent)
}

func (e *Encoding) checkForAlias(o *object) bool {
	for _, i := range e.path {
		for _, z := range i {
			log.Printf("%v not unique %v", o.id, z)
			return true
		}
	}
	return false
}
