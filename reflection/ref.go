package reflection

import (
	"fmt"
	"reflect"
	"strings"
)

func printType(v interface{}) {
	rv := reflect.ValueOf(v)
	typ := rv.Type()
	typeName := ""
	switch rv.Kind() {
	case reflect.Ptr:
		typeName = "pointer"
	case reflect.Int:
		typeName = "int"
	case reflect.Int32:
		typeName = "int32"
	case reflect.String:
		typeName = "string"
	// ... handle more cases
	default:
		typeName = "unrecognized type"
	}
	fmt.Printf("v is of type '%s'. Size: %d bytes\n", typeName, typ.Size())
}

// To minimize API surface, Int() returns int64 and works on all signed integer values (int8, int16, int32, int64).
// All methods for retrieving the value:
// Bool() bool
// Int() int64
// UInt() uint64
// Float() float64
// String() string
// Bytes() []byte
func getIntValue(v interface{}) {
	var reflectValue = reflect.ValueOf(v)
	n := reflectValue.Int()
	fmt.Printf("Int value is: %d\n", n)
}

type S struct {
	N int
}

// Since reflect.ValueOf() creates a reflect.Value that represents a pointer to a value,
// you need to use Elem() to get reflect.Value that represents the value itself.
// You can then call SetInt() to set the value.
func setIntPtr() {
	var n int = 2
	reflect.ValueOf(&n).Elem().SetInt(4)
	fmt.Printf("setIntPtr: n=%d\n", n)
}

func setStructFieldDirect() {
	var s S
	reflect.ValueOf(&s.N).Elem().SetInt(5)
	fmt.Printf("setStructFieldDirect: n=%d\n", s.N)
}

func setStructPtrField() {
	var s S
	reflect.ValueOf(&s).Elem().Field(0).SetInt(6)
	fmt.Printf("setStructPtrField: s.N: %d\n", s.N)
}

func handlePanic(funcName string) {
	if msg := recover(); msg != nil {
		fmt.Printf("%s panicked with '%s'\n", funcName, msg)
	}
}

func setStructField() {
	defer handlePanic("setStructField")
	var s S
	reflect.ValueOf(s).Elem().Field(0).SetInt(4)
	fmt.Printf("s.N: %d\n", s.N)
}

func setInt() {
	defer handlePanic("setInt")
	var n int = 2
	rv := reflect.ValueOf(n)
	rv.Elem().SetInt(4)
}

func setIntPtrWithString() {
	defer handlePanic("setIntPtrWithString")
	var n int = 2
	reflect.ValueOf(&n).Elem().SetString("8")
}

func printIntResolvingPointers(v interface{}) {
	rv := reflect.ValueOf(v)
	typeName := rv.Type().String()
	name := ""
	for rv.Kind() == reflect.Ptr {
		name = "pointer to " + name
		rv = rv.Elem()
	}
	name += rv.Type().String()
	fmt.Printf("Value: %d. Type: '%s' i.e. '%s'.\n\n", rv.Int(), name, typeName)
}

type People struct {
	FirstName string `my_tag:"first-name"`
	lastName  string
	Age       int `json:"age",xml:"AgeXml`
}

func describeStructSimple(rv reflect.Value) {
	structType := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		v := rv.Field(i)
		structField := structType.Field(i)
		name := structField.Name
		typ := structField.Type
		tag := structField.Tag
		jsonTag := tag.Get("json")
		isExported := structField.PkgPath == ""
		if isExported {
			fmt.Printf("name: '%s',\ttype: '%s', value: %v,\ttag: '%s',\tjson tag: '%s'\n", name, typ, v.Interface(), tag, jsonTag)
		} else {
			fmt.Printf("name: '%s',\ttype: '%s',\tvalue: not accessible\n", name, v.Type().Name())
		}
	}
}

func indentStr(level int) string {
	return strings.Repeat("  ", level)
}

// if sf is not nil, this is a field of a struct
func describeStruct(level int, rv reflect.Value, sf *reflect.StructField) {
	structType := rv.Type()
	nFields := rv.NumField()
	typ := rv.Type()
	if sf == nil {
		fmt.Printf("%sstruct %s, %d field(s), size: %d bytes\n", indentStr(level), structType.Name(), nFields, typ.Size())
	} else {
		fmt.Printf("%sname: '%s' type: 'struct %s', offset: %d, %d field(s), size: %d bytes, embedded: %v\n", indentStr(level), sf.Name, structType.Name(), sf.Offset, nFields, typ.Size(), sf.Anonymous)
	}

	for i := 0; i < nFields; i++ {
		fv := rv.Field(i)
		sf := structType.Field(i)
		describeType(level+1, fv, &sf)
	}
}

// if sf is not nil, this is a field of a struct
func describeType(level int, rv reflect.Value, sf *reflect.StructField) {
	switch rv.Kind() {

	case reflect.Int, reflect.Int8:
		// in real code we would handle more primitive types
		i := rv.Int()
		typ := rv.Type()
		if sf == nil {
			fmt.Printf("%stype: '%s', value: '%d'\n", indentStr(level), typ.Name(), i)
		} else {
			fmt.Printf("%s name: '%s' type: '%s', value: '%d', offset: %d, size: %d\n", indentStr(level), sf.Name, typ.Name(), i, sf.Offset, typ.Size())
		}

	case reflect.Ptr:
		fmt.Printf("%spointer\n", indentStr(level))
		describeType(level+1, rv.Elem(), nil)

	case reflect.Struct:
		describeStruct(level, rv, sf)
	}
}

type Inner struct {
	N int
}

type Ss struct {
	Inner
	NamedInner Inner
	PtrInner   *Inner
	unexported int
	N          int8
}

func Run() {
	var v any = 4
	var reflectVal reflect.Value = reflect.ValueOf(v)

	var typ reflect.Type = reflectVal.Type()
	fmt.Printf("Type '%s' of size: %d bytes\n", typ.Name(), typ.Size())
	if typ.Kind() == reflect.Int {
		fmt.Printf("v contains value of type int\n")
	}
	setIntPtr()
	printType(32)
	printType(int64(23))
	printType(int32(12))
	printIntResolvingPointers(3)
	n := 3
	printIntResolvingPointers(&n)
	np := &n
	printIntResolvingPointers(&np)

	s := People{
		FirstName: "John",
		lastName:  "Doe",
		Age:       27,
	}
	describeStructSimple(reflect.ValueOf(s))
	var ss Ss
	describeType(0, reflect.ValueOf(ss), nil)

	// slice example
	a := []int{3, 1, 8}
	rv := reflect.ValueOf(a)

	fmt.Printf("len(a): %d\n", rv.Len())
	fmt.Printf("cap(a): %d\n", rv.Cap())

	fmt.Printf("slice kind: '%s'\n", rv.Kind().String())

	fmt.Printf("element type: '%s'\n", rv.Type().Elem().Name())

	el := rv.Index(0).Interface()
	fmt.Printf("a[0]: %v\n", el)

	elRef := rv.Index(1)
	fmt.Printf("elRef.CanAddr(): %v\n", elRef.CanAddr())
	fmt.Printf("elRef.CanSet(): %v\n", elRef.CanSet())

	elRef.SetInt(5)
	fmt.Printf("a: %v\n", a)
}
