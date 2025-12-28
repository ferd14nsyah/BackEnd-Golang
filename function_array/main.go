package main

import "fmt"

func main() {
	number := [6]int{10, 20, 30, 40, 50, 60}

	fmt.Println("Jumlah elemen :", len(number))
	fmt.Println("Index ke-1 :", number[1])
	number[1] = 90
	fmt.Println("Index ke-1 :", number[1])

	fmt.Println("Ini adalah array :")

	for index, value := range number {
		fmt.Println("isi index ke :", index, " = ", value)
	}

	arr1 := [4]int{1, 3, 5}
	arr2 := [4]int{2, 4, 6}

	fmt.Println("array 1 == array 2 :", arr1 == arr2)
	fmt.Println("array 1 != array 2 :", arr1 != arr2)

}
