package main

import "fmt"

func tambah(a int, b int) int {

	return a - b
	// umur := 12

	// if umur >= 25 {
	// 	fmt.Println("Dewasa")
	// } else {
	// 	fmt.Println("Anak-anak")
	// }
}

func main() {
	hasil := tambah(3, 6)
	fmt.Println("Hasilnya :", hasil)
}
