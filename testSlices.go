package main

import (
	"fmt"
	"reflect"
)

func main() {
	underlyingArray := [3]string{"a", "b", "c"}
	slice2 := underlyingArray[1:]
	fmt.Println(reflect.TypeOf(slice2))
	fmt.Println(reflect.TypeOf(underlyingArray))
}
