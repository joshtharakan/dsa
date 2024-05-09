package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type outExp struct {
	out    []reflect.Value
	expiry time.Time
}

func Cacher(f interface{}, expiry time.Duration) (interface{}, error) {

	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		return nil, errors.New("Cacher only accepts functions")
	}

	// get the input struct
	inputStructType, err := inputStructUsingReflection(ft)
	if err != nil {
		return nil, err
	}

	if ft.NumOut() == 0 {
		fmt.Println("It should have atleast one output param")
		return nil, fmt.Errorf("It should have atleast one output param")
	}

	fmt.Printf("input type looks like this %v\n", inputStructType)

	m := map[interface{}]outExp{}
	fv := reflect.ValueOf(f)

	// write a wrapper function to cache
	cacher := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		// create a reflection struct type value
		inVal := reflect.New(inputStructType).Elem()
		for i, val := range args {
			inVal.Field(i).Set(val)
		}
		inValV := inVal.Interface()
		// check if existing availabl in the map
		ov, ok := m[inValV]
		now := time.Now()
		if !ok || ov.expiry.Before(now) {
			ov.out = fv.Call(args)
			ov.expiry = now.Add(expiry)
			m[inValV] = ov
		}
		return ov.out
	})

	return cacher.Interface(), nil
}

func inputStructUsingReflection(ft reflect.Type) (reflect.Type, error) {
	// checks if the function has input parameters
	if ft.NumIn() == 0 {
		return nil, errors.New("Function has no input parameters")
	}

	var inputStructFields []reflect.StructField

	// convert the inputs into a struct field
	for i := range ft.NumIn() {
		// check if the input is comparable
		if !ft.In(i).Comparable() {
			return nil, fmt.Errorf("Input parameter %d of type %s and kind %v is not comparable", i+1, ft.In(i).Name(), ft.In(i).Kind())
		}
		inputStructFields = append(inputStructFields, reflect.StructField{
			Name: fmt.Sprintf("Field%d", i+1),
			Type: ft.In(i),
		})
	}
	newStruct := reflect.StructOf(inputStructFields)
	return newStruct, nil
}

func AddFn(a, b int) int {
	time.Sleep(200 * time.Millisecond)
	return a + b
}

func main() {

	cache, err := Cacher(AddFn, 2*time.Second)
	if err != nil {
		panic(err)
	}

	cacheeAddFn := cache.(func(int, int) int)
	for _ = range 10 {
		start := time.Now()
		out := cacheeAddFn(1, 2)
		end := time.Now()
		fmt.Printf("got result %v in %v \n", out, end.Sub(start))
	}
	time.Sleep(3 * time.Second)
	start := time.Now()
	out := cacheeAddFn(1, 2)
	end := time.Now()
	fmt.Printf("got result %v in %v \n", out, end.Sub(start))
}
