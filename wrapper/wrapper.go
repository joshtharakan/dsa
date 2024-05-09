package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func wrappedFunction(f interface{}) interface{} {
	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		panic("wrappedFunction only accepts functions")
	}
	fVal := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(fType, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := fVal.Call(in)
		end := time.Now()
		timeTaken := end.Sub(start)
		fmt.Println("Time taken to execute", runtime.FuncForPC(fVal.Pointer()).Name(), "is", timeTaken)
		return out
	})
	return wrapperF.Interface()
}

func timeMe() {
	fmt.Println("Starting Time Me")
	time.Sleep(2 * time.Second)
	fmt.Println("Ending Time Me")
}

func timeMeToo(i int) int {
	fmt.Println("Starting Time Me too")
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println("Ending Time Me")
	return i * 2
}

func main() {
	timed1 := wrappedFunction(timeMe).(func())
	timed1()
	timed2 := wrappedFunction(timeMeToo).(func(int) int)
	fmt.Println(timed2(3))
}
