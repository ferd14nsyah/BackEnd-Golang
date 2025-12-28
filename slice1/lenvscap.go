package main

import "fmt"

func main() {
	//kapasitas(Len Vs Cap)
	buah := []string{"apel", "jeruk", "mangga"}
	fmt.Println("Len:", len(buah)) // jumlah elemen
	fmt.Println("Cap:", cap(buah)) // kapasitas (berapa banyak bisa ditampung sebelum alokasi ulang)

	buah = append(buah, "pisang")
	fmt.Println("Setelah append:", buah)
	fmt.Println("Len:", len(buah), "Cap:", cap(buah))
}
