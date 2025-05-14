package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := "halo, Go, Journey"
	parts := strings.Split(text, ", ")
	fmt.Println(parts)

	upper := strings.ToUpper("golang")
	fmt.Println(upper)

	numStr := "123"
	num, _ := strconv.Atoi(numStr)
	fmt.Println("Angka:", num)

	str := strconv.Itoa(456)
	fmt.Println("String:", str)
}
