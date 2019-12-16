package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:'name'`
	Age  string `json:'age'`
}

func main() {
	var jsonString = `{"name": "Adit", "age": "22"}`
	var jsonData = []byte(jsonString)

	var data User
	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println("Error Unmarshalling Json " + err.Error())
		return
	}

	fmt.Println(data)

	fmt.Println("My name is " + data.Name)
	fmt.Println("My Age is " + data.Age)

}
