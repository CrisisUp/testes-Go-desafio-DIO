package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 10)
	w := subtrai(5, 10)
	z := divide(20, 2)

	fmt.Println(x, y, w, z)
}

func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func subtrai(numeros ...int) int {
	if len(numeros) == 0 {
		return 0
	}
	resultado := numeros[0]
	for _, num := range numeros[1:] {
		resultado -= num
	}
	return resultado
}

func multiplica(numeros ...int) int {
	total := 1
	for _, num := range numeros {
		total *= num
	}
	return total
}

func divide(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0
	}
	resultado := numeros[0]
	for _, num := range numeros[1:] {
		resultado /= num
	}
	return resultado
}
