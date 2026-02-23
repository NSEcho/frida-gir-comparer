package types

import "C"
import (
	"fmt"
)

// GoTypes is a basic mapping between C types and Go types
var GoTypes = map[string]string{
	"guint":  "uint",
	"gint":   "int",
	"gchar*": "string",
}

/*
TODO: instead of writing this boilerplate code make some of them being
parsed in function call
*/
type wrapFnParam func(paramName string) (string, string)

var gcharRet = `%sC := C.CString(%s)
	defer C.free(unsafe.Pointer(%sC))`

func paramGCharPointer(paramName string) (string, string) {
	desc := fmt.Sprintf(gcharRet, paramName, paramName, paramName)
	return desc, paramName + "C"
}

func paramIntToInt(paramName string) (string, string) {
	desc := fmt.Sprintf(`%sC := C.int(%s)`, paramName, paramName)
	return desc, paramName + "C"
}

func paramGuintToUint(paramName string) (string, string) {
	desc := fmt.Sprintf(`%sC := C.uint(%s)`, paramName, paramName)
	return desc, paramName + "C"
}

var FnParams = map[string]wrapFnParam{
	"gchar*": paramGCharPointer,
	"int":    paramIntToInt,
	"guint":  paramGuintToUint,
}

type retTypeFn func() string

func gcharTypeToString() string {
	return "C.GoString(rt)"
}

func guintToUint() string {
	return "uint(rt)"
}

func intToInt() string {
	return "int(rt)"
}

var RetTypesFns = map[string]retTypeFn{
	"string": gcharTypeToString,
	"uint":   guintToUint,
	"int":    intToInt,
}
