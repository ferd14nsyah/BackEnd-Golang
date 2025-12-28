package main

import (
	"fmt"
	"strconv"
)

func main() {
	//bool -> string

	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean ke String:", str)

	//string -> Bool
	val, _ := strconv.ParseBool("true")
	fmt.Println("String ke Boolean:", val)
}
