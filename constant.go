package main

import "fmt"

// constant adalah variable yang tidak dapat di ubah lagi setelah di beri nilai
func main() {
	const firstName = "Bima"
	const lastName = "Bima Bagaskhoro"
	const value = 1000

	fmt.Println(firstName, lastName, value)

	// Deklarasi Multiple Constant

	const (
		firstNames = "Bima"
		lastNames  = "Bagaskhoro"
		age        = 22
	)

	fmt.Println(firstNames, lastNames, age)
}
