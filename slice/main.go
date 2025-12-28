package main

import "fmt"

func main() {
	angka := [5]int{1, 2, 3, 4, 5}
	slice := angka[1:4] // ambil index 1 sampai 3 (2,3,4)

	fmt.Println("Slice:", slice)
	slice[0] = 99                // ubah index 1 pada array
	fmt.Println("Array:", angka) // ikut berubah!
}
