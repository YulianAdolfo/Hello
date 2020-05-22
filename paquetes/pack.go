package main

import (
	"fmt"

	"./greet"
)

func main() {
	fmt.Println(greet.InEnglish("Yulian"))
	fmt.Println(greet.InSpanish("Yulian"))
	fmt.Println(greet.InItalian("Yulian"))
}
