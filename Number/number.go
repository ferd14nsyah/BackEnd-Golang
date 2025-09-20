package main

import "fmt"

func main() {

	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 92311298130982132

	var i int = -100

	var f32 float32 = 33.231213123
	var f64 float64 = 172837187239817291

	var u8 uint8 = 233
	var u16 uint16 = 63276
	var u32 uint32 = 3218901830
	var u64 uint64 = 127387128373

	var u uint = 100

	fmt.Println("Signed Integers:")
	fmt.Printf("int8 : %v\n", i8)
	fmt.Printf("int16 : %v\n", i16)
	fmt.Printf("int32 : %v\n", i32)
	fmt.Printf("int64 : %v\n", i64)
	fmt.Printf("i : %v\n", i)

	fmt.Println("\nUnSigned Integers:")
	fmt.Printf("uint8 : %v\n", u8)
	fmt.Printf("unt16 : %v\n", u16)
	fmt.Printf("unt32 : %v\n", u32)
	fmt.Printf("unt64 : %v\n", u64)
	fmt.Printf("u : %v\n", u)

	fmt.Println("\nNilai dari Float32", f32)
	fmt.Println("Nilai dari Float64", f64)

}
