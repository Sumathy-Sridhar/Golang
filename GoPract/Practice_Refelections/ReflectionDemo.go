package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	cusname string
	mobile  int
}

type Cake struct {
	cakename string
	quantity int
	amount   float32
}

func insertQuery(ck interface{}) {
	typechk := reflect.TypeOf(ck) //returns the type of struct
	value := reflect.ValueOf(ck)  //returns value of type
	kind := typechk.Kind()        // returns if it is a struct or any type
	fmt.Println("Type : ", typechk)
	fmt.Println("Values : ", value)
	fmt.Println("Number of fields: ", value.NumField())
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field:%d type:%T value:%v\n", i, value.Field(i), value.Field(i))
	}
	fmt.Println("Kind: ", kind)
}

func main() {
	// cus := customer{
	// 	cusname: "Sumathy",
	// 	mobile:  992587469,
	// }
	// insertQuery(cus)
	cakevar := Cake{
		cakename: "ChocoTruffle",
		quantity: 1,
		amount:   500.25,
	}
	insertQuery(cakevar)
}
