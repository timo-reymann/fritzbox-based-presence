package main

import (
	"fmt"
	"reflect"
)

type test struct {
	foo string
}

func main1() {
	// one way is to have a value of the type you want already
	a := test{}
	// reflect.New works kind of like the built-in function new
	// We'll get a reflected pointer to a new int value
	newInstance := reflect.New(reflect.TypeOf(a))
	// Just to prove it
	b := newInstance.Elem()
	b.Set(reflect.ValueOf(test{"test"}))
	// Prints 0
	fmt.Println(b)
}
