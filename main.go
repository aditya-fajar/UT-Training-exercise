package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string `json:'name'`
	Age  string `json:'age'`
}

func tes() {
	//JSON-------------------------------------------------------------------------------
	// var jsonString = `{"name": "Adit", "age": "22"}`
	// var jsonData = []byte(jsonString)

	// var data User
	// var err = json.Unmarshal(jsonData, &data)

	// if err != nil {
	// 	fmt.Println("Error Unmarshalling Json " + err.Error())
	// 	return
	// }

	// fmt.Println(data)

	// fmt.Println("My name is " + data.Name)
	// fmt.Println("My Age is " + data.Age)

	//SLICE----------------------------------------------------------------------------
	// var slice = []string{"a", "b", "c", "d"}

	// var slice2 = slice[1:2]

	// fmt.Println(slice)
	// fmt.Println(slice2)

	// slice2 = []string{"1", "2", "3"}

	// fmt.Println(slice)
	// fmt.Println(slice2)

	//CALC-----------------------------------------------------------------------------------
	// number := 8.0

	// result, err := calculate(number)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	//POINTER----------------------------------------------------------------------------------
	var a int = 4
	var b *int = &a

	fmt.Println("Number a {value} 	: ", a)
	fmt.Println("Number a {address} : ", &a)
	fmt.Println("Number b {value} 	: ", *b)
	fmt.Println("Number b {address} : ", b)

	fmt.Println("===================================")

	a = 5
	fmt.Println("Number a {value} 	: ", a)
	fmt.Println("Number a {address} : ", &a)
	fmt.Println("Number b {value} 	: ", *b)
	fmt.Println("Number b {address} : ", b)

	fmt.Println("===================================")

	*b = 6
	fmt.Println("Number a {value} 	: ", a)
	fmt.Println("Number a {address} : ", &a)
	fmt.Println("Number b {value} 	: ", *b)
	fmt.Println("Number b {address} : ", b)

}

func calculate(d float64) (float64, error) {
	if d == 8 {
		return 0, errors.New("don't 8")
	}

	cal := d + 10

	return cal, nil
}
