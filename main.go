package main

import "fmt"

func main() {

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Hallo belajar golang : ", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Hallo belajar golang : ", i)
	// 	i++
	// }

	// title := "Golang Fazil"

	// for _, letter := range title {
	// 	fmt.Println("letter :", string(letter))
	// }

	// var bahasas [5]string
	// bahasas[0] = "F"
	// bahasas[1] = "A"
	// bahasas[2] = "Z"
	// bahasas[3] = "I"
	// bahasas[4] = "L"

	bahasas := [...]string{
		"F",
		"A",
		"Z",
		"I",
		"L",
	}

	// fmt.Println(bahasas)
	// length := len(bahasas)
	// fmt.Println(length)

	for _, baha := range bahasas {
		fmt.Println(baha)
	}
}
