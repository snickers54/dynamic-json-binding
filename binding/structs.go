package binding

import (
	"fmt"
	"go/ast"
	"log"
	"regexp"
	"sort"
	"strings"
)

func saveStruct(name string, t *ast.StructType) {
	signature := []string{}
	jsonTagRegex := regexp.MustCompile("json:\"[a-zA-Z0-9]+\"")
	for _, list := range t.Fields.List {
		name := getJSONTag(list.Tag, jsonTagRegex)
		if len(name) == 0 {
			name = list.Names[0].Name
		}
		signature = append(signature, strings.ToLower(name))
	}
	bindingHelper.Mutex.Lock()
	if isDuplicated, key := isDuplicate(signature); isDuplicated == true {
		log.Fatalf("Unfortunately, these two structs (%s, %s) signatures are identitcal.\n Use json tags to give aliases to fields of your struct and implicitly modify your struct signature.", key, name)
	}
	bindingHelper.SimplifiedStructs[currentPkg+"."+name] = signature
	bindingHelper.Mutex.Unlock()
}

func getJSONTag(tag *ast.BasicLit, jsonTagRegex *regexp.Regexp) string {
	if tag == nil {
		return ""
	}
	foundTag := strings.Replace(jsonTagRegex.FindString(strings.Replace(tag.Value, " ", "", -1)), "\"", "", -1)
	return strings.Replace(foundTag, "json:", "", -1)
}

func isDuplicate(newSlice sort.StringSlice) (bool, string) {
	newSlice.Sort()
	for key, stringSlice := range bindingHelper.SimplifiedStructs {
		stringSlice.Sort()
		if isSliceEqual(stringSlice, newSlice) {
			return true, key
		}
	}
	return false, ""
}

func isSliceEqual(a, b sort.StringSlice) bool {
	return fmt.Sprintf("%+v", a) == fmt.Sprintf("%+v", b)
}

func findStructTypes(stringSlice sort.StringSlice) []string {
	results := []string{}
	stringSlice.Sort()
	for key, slice := range bindingHelper.SimplifiedStructs {
		_, exists := bindingHelper.MapTypes[key]
		if exists == false {
			continue
		}
		slice.Sort()
		if isSliceEqual(stringSlice, slice) {
			results = append(results, key)
		}
	}
	return results
}
