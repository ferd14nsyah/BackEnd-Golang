package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hallo My Name Is Ferdian"
	var isRaining bool = true
	var isSunny bool = false

	isLoggedId := true
	haspermission := false

	fmt.Println(`Apakah Hujan :`, isRaining)
	fmt.Println(`Apakah Cuacanya :`, isSunny)
	fmt.Println(`Apakah Kamu Akan Login :`, isLoggedId)
	fmt.Println(`Apakah Kamu Punya Autorized :`, haspermission)

	fmt.Println("Huruf Lowercase", strings.ToLower(text))
	fmt.Println("Huruf Uppercase", strings.ToUpper(text))
	fmt.Println("StartWith Halo?", strings.HasPrefix(text, "Hallo"))
	fmt.Println("Contains My Name?", strings.Contains(text, "My Name"))

	parts := strings.Split("Apel, Banana Asem", ",")
	fmt.Println("Split:", parts)

	newtext := strings.ReplaceAll(text, "Name", "Ganteng")
	fmt.Println("Replace", newtext)
}
