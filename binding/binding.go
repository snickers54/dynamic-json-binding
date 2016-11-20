package binding

import (
	"encoding/json"
	"errors"
	"go/ast"
	"reflect"
	"sort"
	"strings"
	"sync"
)

var bindingHelper BindingHelper = BindingHelper{}

type BindingHelper struct {
	GoFiles           []string
	SimplifiedStructs map[string]sort.StringSlice
	Mutex             sync.Mutex
	MapTypes          map[string]interface{}
}

func init() {
	listFS("./")
	bindingHelper.MapTypes = map[string]interface{}{}
	bindingHelper.SimplifiedStructs = map[string]sort.StringSlice{}
	for _, path := range bindingHelper.GoFiles {
		getASTFromFile(path)
	}
}

func (self *BindingHelper) Visit(node ast.Node) (w ast.Visitor) {
	switch t := node.(type) {
	case *ast.TypeSpec:
		switch t1 := t.Type.(type) {
		case *ast.StructType:
			saveStruct(t.Name.Name, t1)
		}
	}
	return self
}

func RegisterTypes(types ...interface{}) {
	for _, t := range types {
		bindingHelper.MapTypes[reflect.TypeOf(t).String()] = t
	}
}

func Bind(content []byte) (interface{}, error) {
	value := map[string]json.RawMessage{}
	err := json.Unmarshal(content, &value)
	if err != nil {
		return nil, err
	}
	stringSlice := extractKeys(value)
	structTypes := findStructTypes(stringSlice)
	if len(structTypes) > 1 || len(structTypes) == 0 {
		return nil, errors.New("Zero or more than one matching struct found.")
	}
	t := reflect.TypeOf(bindingHelper.MapTypes[structTypes[0]])
	t1 := reflect.New(t).Interface()
	err = json.Unmarshal(content, &t1)
	if err != nil {
		return nil, err
	}
	return t1, nil
}

func extractKeys(value map[string]json.RawMessage) sort.StringSlice {
	stringSlice := sort.StringSlice{}
	for key, _ := range value {
		stringSlice = append(stringSlice, strings.ToLower(key))
	}
	return stringSlice
}
