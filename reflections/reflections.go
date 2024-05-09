package main

import (
	"fmt"
	"reflect"
	"strings"
)

type SampleStruct struct {
	Name string
	Age  int
}

func evaluate(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "contained type of", t.Kind())
		evaluate(t.Elem(), depth+1)
	case reflect.Struct:
		for i := range t.NumField() {
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "is", t.Field(i).Name, "of type", t.Field(i).Type.Name(), "and kind", t.Field(i).Type.Kind())
		}

	}
}

func makeStruct(fields ...interface{}) interface{} {
	sfs := make([]reflect.StructField, len(fields))
	for i, f := range fields {
		sfs[i] = reflect.StructField{
			Name: fmt.Sprintf("Field%d", i+1),
			Type: reflect.TypeOf(f),
		}
	}
	st := reflect.StructOf(sfs)
	so := reflect.New(st)
	return so.Interface()
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	mp := map[string]int{"one": 1, "two": 2, "three": 3}
	greeting := "Hello, World!"
	greetingP := &greeting
	st := SampleStruct{
		Name: "John Doe",
		Age:  30,
	}
	stP := &st

	slType := reflect.TypeOf(sl)
	mpType := reflect.TypeOf(mp)
	greetingType := reflect.TypeOf(greeting)
	greetingPType := reflect.TypeOf(greetingP)
	stType := reflect.TypeOf(st)
	stPType := reflect.TypeOf(stP)

	evaluate(slType, 0)
	evaluate(mpType, 0)
	evaluate(greetingType, 0)
	evaluate(greetingPType, 0)
	evaluate(stType, 0)
	evaluate(stPType, 0)

	// view content of a reflection
	fmt.Println("Value of sl is", reflect.ValueOf(sl))
	fmt.Println("Value of mp is", reflect.ValueOf(mp))
	fmt.Println("Value of greeting is", reflect.ValueOf(greeting))
	fmt.Println("Value of greetingP is", reflect.ValueOf(greetingP).Elem()) // dereference the pointer
	fmt.Println("Value of st is", reflect.ValueOf(st))
	fmt.Println("Value of stP is", reflect.ValueOf(stP).Elem()) // dereference the pointer

	// change the value of a reflection
	reflect.ValueOf(&greeting).Elem().SetString("Hello, Universe!")
	fmt.Println("updated value of greeting is", reflect.ValueOf(greeting))

	// create a new reflection
	newStruct := reflect.New(stType)
	newStruct.Elem().Field(0).SetString("Jane Doe")
	newStruct.Elem().Field(1).SetInt(25)
	fmt.Println("new struct reflection is", newStruct.Elem())

	st2 := newStruct.Elem().Interface().(SampleStruct)
	fmt.Printf("new struct is %v, name is %s, age is %v\n", st2, st2.Name, st2.Age)

	// create a new slice reflection
	sliceR := reflect.MakeSlice(slType, 0, 0)
	sliceR = reflect.Append(sliceR, reflect.ValueOf(6))
	sliceR = reflect.Append(sliceR, reflect.ValueOf(7))
	newSlice := sliceR.Interface().([]int)
	fmt.Println("new slice is", newSlice)

	// create a new map reflection
	mapReflect := reflect.MakeMap(mpType)
	mapReflect.SetMapIndex(reflect.ValueOf("four"), reflect.ValueOf(4))
	mapReflect.SetMapIndex(reflect.ValueOf("five"), reflect.ValueOf(5))
	newMap := mapReflect.Interface().(map[string]int)
	fmt.Println("new map is", newMap)

	fmt.Println("Creating a struct with fields of different types")
	// create a struct with based on given array of fields
	newStruct2 := makeStruct(1, "two", []int{3.0, 4.0})
	// this returned a pointer to a struct with 3 fields:
	// an int, a string, and a slice of ints
	// but you can’t actually use any of these fields
	// directly in the code; you have to reflect them
	sr := reflect.ValueOf(newStruct2)

	// getting and setting the int field
	fmt.Println("current value for field 0", sr.Elem().Field(0).Interface())
	sr.Elem().Field(0).SetInt(20)
	fmt.Println("updated value for field 0", sr.Elem().Field(0).Interface())

	// getting and setting the string field
	fmt.Println("current value for field 1", sr.Elem().Field(1).Interface())
	sr.Elem().Field(1).SetString("reflect me")
	fmt.Println("updated value for field 1", sr.Elem().Field(1).Interface())

	// getting and setting the []int field
	fmt.Println("current value for field 2", sr.Elem().Field(2).Interface())
	v := []int{1, 2, 3}
	rv := reflect.ValueOf(v)
	sr.Elem().Field(2).Set(rv)
	fmt.Println("updated value for field 2", sr.Elem().Field(2).Interface())
}
