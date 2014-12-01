package gen

import (
	"io"
	"strconv"
	"strings"
	"text/template"
)

// genCallbacks creates the C wrapper functions for the callbacks and writes a go
// file with the wrappers to the given io.Writer. There is an accompanying test that
// when run creates the wrapper function in the standard src directory.
func genCallbacks(count int, w io.Writer) {
	t, _ := template.New("callback generator").Parse(callbackTemplate)
	paramLists := [][]paramType{
		[]paramType{},
		[]paramType{intType},
		[]paramType{intType, intType},
		[]paramType{intType, intType, stringType},
		[]paramType{intType, intType, intType, intType, stringType},
		[]paramType{stringType, intType, intType, intType},
	}
	t.Execute(w, struct {
		Externs     string
		Exports     string
		CCallbacks  string
		GoCallbacks string
		Checks      string
	}{
		genAllExterns(count, paramLists),
		genExports(count, paramLists),
		genCCallbackArrays(count, paramLists),
		genGoCallbackArrayVars(paramLists),
		genFuncTypeChecks(paramLists),
	})
}

const callbackTemplate = `package gen

import "reflect"

/*

#cgo linux LDFLAGS: -liup
#include <iup.h>
#include <stdlib.h>

{{.Externs}}
*/
import "C"

{{.Exports}}{{.CCallbacks}}{{.GoCallbacks}}
func GetNextCallback(cb interface{}) C.Icallback {
	typ := reflect.TypeOf(cb)
	checkIntReturningFunction(typ)
{{.Checks}}	panic("unsupported callback function type")
}

func checkIntReturningFunction(typ reflect.Type) {
	if typ.Kind() != reflect.Func {
		panic("callback must be a function")
	}
	if typ.NumOut() != 1 || typ.Out(0).String() != "int" {
		panic("callback must return int")
	}
}

func checkEnoughCallbacks(inGo, inC int) {
	if inGo > inC-1 {
		panic("not enough callbacks, regenerate more")
	}
}
`

type paramType struct {
	funcName     string
	cType        string
	cTypeInGo    string
	goType       string
	goConversion string
}

var intType = paramType{"int", "int", "C.int", "int", "int"}
var stringType = paramType{"string", "char*", "*C.char", "string", "C.GoString"}

func genAllExterns(count int, typeLists [][]paramType) string {
	externs := ""
	for _, types := range typeLists {
		externs += genExterns(count, types)
	}
	return externs
}

// multiple lines à la:
// extern int go_cb_handle_int_string_25(Ihandle* h, int arg0, const char* arg1);
func genExterns(count int, types []paramType) string {
	funcName := makeCFuncName(types)
	params := makeCParams(types)
	externs := ""
	for i := 0; i < count; i++ {
		externs += "extern int " + funcName + strconv.Itoa(i) + params + ";\n"
	}
	return externs
}

// go_cb_handle_int_string_int_
func makeCFuncName(types []paramType) string {
	return "go_cb_handle_" + makeFuncNameTypes(types) + "_"
}

// int_string_int
func makeFuncNameTypes(types []paramType) string {
	if len(types) == 0 {
		return "none"
	}
	typeNames := make([]string, len(types))
	for i, t := range types {
		typeNames[i] = t.funcName
	}
	return strings.Join(typeNames, "_")
}

// (Ihandle* h, int arg0, const char* arg1)
func makeCParams(types []paramType) string {
	s := "(Ihandle* h"
	for i, t := range types {
		s += ", " + t.cType + " arg" + strconv.Itoa(i)
	}
	return s + ")"
}

func genExports(count int, typeLists [][]paramType) string {
	exports := ""
	for _, types := range typeLists {
		exports += genExport(count, types)
	}
	return exports
}

func genExport(count int, types []paramType) string {
	funcNameStart := makeCFuncName(types)
	params := makeCParamsInGo(types)
	typeNames := makeFuncNameTypes(types)
	exports := ""
	for i := 0; i < count; i++ {
		index := strconv.Itoa(i)
		funcName := funcNameStart + index
		exports += "//export " + funcName + "\n"
		exports += "func " + funcName + params +
			" C.int { return C.int(callbacks_" + typeNames + "[" + index + "](" +
			castArgsToGo(types) + ")) }\n\n"
	}
	return exports
}

// (h *C.Ihandle, arg0 int, arg1 *C.char)
func makeCParamsInGo(types []paramType) string {
	s := "(h *C.Ihandle"
	for i, t := range types {
		s += ", arg" + strconv.Itoa(i) + " " + t.cTypeInGo
	}
	return s + ")"
}

func castArgsToGo(types []paramType) string {
	cast := make([]string, len(types))
	for i, t := range types {
		cast[i] = t.goConversion + "(arg" + strconv.Itoa(i) + ")"
	}
	return strings.Join(cast, ", ")
}

func genCCallbackArrays(count int, types [][]paramType) string {
	s := ""
	for _, t := range types {
		s += genCCallbackArray(count, t)
	}
	return s
}

func genCCallbackArray(count int, types []paramType) string {
	funcName := makeCFuncName(types)
	s := "var cCallbacks_" + makeFuncNameTypes(types) + " = []C.Icallback{\n"
	for i := 0; i < count; i++ {
		s += "\t(C.Icallback)(C." + funcName + strconv.Itoa(i) + "),\n"
	}
	s += "}\n\n"
	return s
}

func genGoCallbackArrayVars(types [][]paramType) string {
	s := ""
	for _, t := range types {
		s += genGoCallbackArrayVar(t)
	}
	return s
}

func genGoCallbackArrayVar(types []paramType) string {
	return "var callbacks_" + makeFuncNameTypes(types) + " []func" + makeGoParams(types) + " int\n"
}

// (arg0 int, arg1 string)
func makeGoParams(types []paramType) string {
	args := make([]string, len(types))
	for i, t := range types {
		args[i] = "arg" + strconv.Itoa(i) + " " + t.goType
	}
	return "(" + strings.Join(args, ", ") + ")"
}

func genFuncTypeChecks(types [][]paramType) string {
	s := ""
	for _, t := range types {
		s += genFuncTypeCheck(t)
	}
	return s
}

func genFuncTypeCheck(types []paramType) string {
	s := "\tif typ.NumIn() == " + strconv.Itoa(len(types)) + reflectTypeChecks(types) + " {\n"
	name := makeFuncNameTypes(types)
	s += "\t\tcheckEnoughCallbacks(len(callbacks_" + name + "), len(cCallbacks_" + name + "))\n"
	s += "\t\tcallbacks_" + name + " = append(callbacks_" + name + ", cb.(func(" + goTypeOnlyList(types) + ") int))\n"
	s += "\t\treturn cCallbacks_" + name + "[len(callbacks_" + name + ")-1]\n"
	s += "\t}\n"
	return s
}

func reflectTypeChecks(types []paramType) string {
	s := ""
	for i, t := range types {
		s += " && typ.In(" + strconv.Itoa(i) + ").String() == \"" + t.goType + "\""
	}
	return s
}

// int, string, int
func goTypeOnlyList(types []paramType) string {
	names := make([]string, len(types))
	for i, t := range types {
		names[i] = t.goType
	}
	return strings.Join(names, ", ")
}
