package main

import "fmt"

func main() {
	x := 8.0
	y := 10.0

	tambah := add(x, y)
	kurang := min(x, y)
	kali := multiple(x, y)
	bagi := division(x, y)

	fmt.Println("Penjumlahan", x, " + ", y, " = ", tambah)
	fmt.Println("Pengurangan", x, " - ", y, " = ", kurang)
	fmt.Println("Perkalian", x, " * ", y, " = ", kali)
	fmt.Println("Pembagian", x, " / ", y, " = ", bagi)
}

func add(x float64, y float64) float64 {
	add := x + y
	return add
}

func min(x float64, y float64) float64 {
	min := x - y
	return min
}

func multiple(x float64, y float64) float64 {
	multiple := x * y
	return multiple
}

func division(x float64, y float64) float64 {
	division := x / y
	return division
}
