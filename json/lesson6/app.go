package main

import (
	"reflect"
	"fmt"
)

func main() {

	type t struct {
		Name string
		N int
	}
	var n = t{"John",42}
	ps := reflect.ValueOf(&n)
	s := ps.Elem()
	s.Set()

	// N at start
	//fmt.Println(n.N)
	//// pointer to struct - addressable
	//ps := reflect.ValueOf(&n)
	//// struct
	//s := ps.Elem()
	//if s.Kind() == reflect.Struct {
	//	// exported field
	//	f := s.FieldByName("N")
	//	if f.IsValid() {
	//		// A Value can be changed only if it is
	//		// addressable and was not obtained by
	//		// the use of unexported struct fields.
	//		if f.CanSet() {
	//			// change value of N
	//			if f.Kind() == reflect.Int {
	//				x := int64(7)
	//				if !f.OverflowInt(x) {
	//					f.SetInt(x)
	//				}
	//			}
	//		}
	//	}
	//}
	//// N at end





}
