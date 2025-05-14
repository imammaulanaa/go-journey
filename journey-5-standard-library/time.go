package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Sekarang:", now)

	oneHourLater := now.Add(time.Hour)
	fmt.Println("Satu jam lagi:", oneHourLater)

	formatted := now.Format("02-Jan-2006 15:04:05")
	fmt.Println("Formatted time:", formatted)

	time.Sleep(2 * time.Second)
	fmt.Println("2 detik berlalu...")
}
