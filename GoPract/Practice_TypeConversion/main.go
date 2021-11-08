package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	//convert string to integer
	stringvar := "100"
	intvar, err := strconv.Atoi(stringvar) //Atoi()  [func Atoi(s string) (int,error)]
	fmt.Println(intvar, err, reflect.TypeOf(intvar))
	svar := "1000"
	pintvar, perr := strconv.ParseInt(svar, 0, 16) // PArseInt()
	fmt.Println(pintvar, perr, reflect.TypeOf(pintvar))
	strVar := "100"
	intValue := 0
	_, eerr := fmt.Sscan(strVar, &intValue) // Sscan()
	fmt.Println(intValue, eerr, reflect.TypeOf(intValue))

	//convert string to Float
	s := "3.1415926535"
	f, err := strconv.ParseFloat(s, 8) //ParseFloat
	fmt.Println(f, err, reflect.TypeOf(f))

	//Convert String to Bool
	s1 := "true"
	b1, _ := strconv.ParseBool(s1)
	fmt.Printf("%T, %v\n", b1, b1)
	s3 := "0"
	b3, _ := strconv.ParseBool(s3) //ParseBool()
	fmt.Printf("%T, %v\n", b3, b3)

	//convert float to string
	var fl float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(fl))
	fmt.Println(fl)
	var st string = strconv.FormatFloat(f, 'E', -1, 32) //FormatFloat()
	fmt.Println(reflect.TypeOf(st))
	fmt.Println(st)

	b := 12.454
	fmt.Println(reflect.TypeOf(b))

	si := fmt.Sprintf("%v", b) //Sprintf()
	fmt.Println(si)
	fmt.Println(reflect.TypeOf(si))

	//convert int to string
	var i int64 = 125
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)

	var str string = strconv.FormatInt(i, 10) //FormatInt()
	fmt.Println(reflect.TypeOf(str))

	fmt.Println("Base 10 value of s:", str)

	//convert int to int16 int32 int 64
	var val int = 10
	fmt.Println(reflect.TypeOf(val))
	val16 := int16(val)
	fmt.Println(val16)
	val32 := int32(val)
	fmt.Println(val32)

	//Convert Float32 to Float64 and Float64 to Float32
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	f64 := float64(f32)
	fmt.Println(reflect.TypeOf(f64))

	//convert int to float
	var fl32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(fl32))

	i32 := int32(fl32)
	fmt.Println(reflect.TypeOf(i32))
	fmt.Println(i32)

}
