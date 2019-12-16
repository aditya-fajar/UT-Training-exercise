package main

import (
	"fmt"
)

func main() {
	var x, y, calculate float32
	var s string
	fmt.Printf("Masukkan Nilai 1 : ")
	fmt.Scanln(&x)

	fmt.Printf("Masukkan Modifier (+, -, *, /) : ")
	fmt.Scanln(&s)

	fmt.Printf("Masukkan Nilai 2 : ")
	fmt.Scanln(&y)

	calculate = arithmatic(x, y, s)

	fmt.Println("Hasil Aritmatik ", x, s, y, "=", calculate)
}

func arithmatic(x float32, y float32, s string) float32 {
	switch s {
	case "+":
		return x + y
		break

	case "-":
		return x - y
		break
	case "*":
		return x * y
		break
	case "/":
		return x / y
		break
	}

	return 0

}
