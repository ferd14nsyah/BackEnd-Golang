package main

import "fmt"

func main() {
	// Array
	var angka = [6]int{10, 20, 30, 40, 50, 60}

	// Slice (lebih fleksibel)
	buah := []string{"apel", "mangga", "pisang"}

	fmt.Println("Array:", angka)
	fmt.Println("Slice:", buah[2])

	S1 := angka[:]
	fmt.Println("Angka dari S1 :", S1)

	S2 := angka[:3]
	fmt.Println("Angka dari S2 :", S2)

	S3 := angka[2:]
	fmt.Println("Angka dari S3 :", S3)

	S4 := angka[1:5]
	fmt.Println("Angka dari S4 :", S4)
}
