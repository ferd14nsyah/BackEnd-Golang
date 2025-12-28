package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}

	s1 := arr[:]
	fmt.Println("Ini adalah nilai dari S1 :", s1)

	s2 := arr[:2]
	fmt.Println("Ini adalah Nilai dari S2 :", s2)

	s3 := arr[4:]
	fmt.Println("Ini adalah Nilai dari S3 :", s3)

	s4 := arr[2:5]
	fmt.Println("Ini adalah Nilai dari S4 :", s4)
}
