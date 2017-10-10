package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
)

const (
	Uint8      = "Uint8"
	Uint16     = "Uint16"
	Uint32     = "Uint32"
	Uint64     = "Uint64"
	Int8       = "Int8"
	Int16      = "Int16"
	Int32      = "Int32"
	Int64      = "Int64"
	Float32    = "Float32"
	Float64    = "Float64"
	Complex64  = "Complex64"
	Complex128 = "Complex128"
	Byte       = "Byte"
	Rune       = "Rune"
	Uint       = "Uint"
	Int        = "Int"
	Uintptr    = "Uintptr"
	String     = "String"
	Interface  = "Interface"
)

var types = []string{
	Interface,
	Uint8,
	Uint16,
	Uint32,
	Uint64,
	Int8,
	Int16,
	Int32,
	Int64,
	Float32,
	Float64,
	Complex64,
	Complex128,
	Byte,
	Rune,
	Uint,
	Int,
	Uintptr,
	String,
}

var typesMap = map[string]string{
	Uint8:      "uint8",
	Uint16:     "uint16",
	Uint32:     "uint32",
	Uint64:     "uint64",
	Int8:       "int8",
	Int16:      "int16",
	Int32:      "int32",
	Int64:      "int64",
	Float32:    "float32",
	Float64:    "float64",
	Complex64:  "complex64",
	Complex128: "complex128",
	Byte:       "byte",
	Rune:       "rune",
	Uint:       "uint",
	Int:        "int",
	Uintptr:    "uintptr",
	String:     "string",
	Interface:  "interface{}",
}

func main() {
	GenCode()
}

func GenCode() {
	var buf bytes.Buffer
	buf.WriteString("package IF\n\n")

	// Else function
	for _, key := range types {
		t := typesMap[key]
		if key == Interface {
			fmt.Fprintf(&buf,
				"func Else(condition bool, a %s, b %s) %s {\n",
				t, t, t,
			)
			buf.WriteString("if condition {\nreturn a\n}\nreturn b\n}\n\n")
			fmt.Fprintf(&buf,
				"func ElseSlice(condition bool, a []%s, b []%s) []%s {\n",
				t, t, t,
			)
			buf.WriteString("if condition {\nreturn a\n}\nreturn b\n}\n\n")

			// Nil function
			fmt.Fprintf(&buf,
				"func Nil(object %s, a %s) %s {\n",
				t, t, t,
			)
			buf.WriteString("if isNil(object) {\nreturn a}\nreturn object\n}\n\n")

			// NotNil function
			fmt.Fprintf(&buf,
				"func NotNil(object %s, a %s) %s {\n",
				t, t, t,
			)
			buf.WriteString("if !isNil(object) {\nreturn a}\nreturn object\n}\n\n")
		} else {
			fmt.Fprintf(&buf,
				"func Else%s(condition bool, a %s, b %s) %s {\n",
				key, t, t, t,
			)
			buf.WriteString("if condition {\nreturn a\n}\nreturn b\n}\n\n")
			fmt.Fprintf(&buf,
				"func Else%sSlice(condition bool, a []%s, b []%s) []%s {\n",
				key, t, t, t,
			)
			buf.WriteString("if condition {\nreturn a\n}\nreturn b\n}\n\n")

			// Empty
			switch key {
			case Uint8, Uint16, Uint32, Uint64, Int8, Int16, Int32, Int64,
				Float32, Float64, Complex64, Complex128, Rune, Byte, Uint, Int:
				fmt.Fprintf(&buf,
					"func Zero%s(n %s, a %s) %s {\n",
					key, t, t, t,
				)
				fmt.Fprintf(&buf, "if n == %s(0) {\nreturn a\n}\nreturn n\n}\n\n", t)
			case String:
				fmt.Fprintf(&buf,
					"func Empty(str %s, a %s) %s {\n",
					t, t, t,
				)
				buf.WriteString(`if str == "" {
return str
}
return str
}

`)
			}
		}
	}

	buf.WriteString(`// isNil is stolen from https://github.com/stretchr/testify/blob/master/assert/assertions.go
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	if reflect.Chan <= kind && kind <= reflect.Slice && value.IsNil() {
		return true
	}
	return false
}`)

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	f, err := os.Create("if.go")
	if err != nil {
		log.Fatal(err)
	}
	f.Write(formatted)
	f.Close()
}
