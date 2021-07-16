package main

import "fmt"

func main() {

	var pembatas = "---------------------------------------------------"

	var name string
	var status string
	var jurusan string
	var university string
	var age int

	name = "Bima Bagaskhoro"
	status = "mahasiswa"
	jurusan = "Fakultas Ilmu Komputer"
	university = "Bhayangkara Jakarta Raya University"
	age = 22

	fmt.Println(pembatas)

	fmt.Println(name)
	fmt.Println(status)
	fmt.Println(name, status, jurusan, university, age)

	fmt.Println(pembatas)

	//metode lainnya
	var names = "Bima Bagaskhoro"
	var hobby = "Learn"
	var ages = 22
	fmt.Println(names, hobby, ages)

	fmt.Println(pembatas)

	//metode kata kunci var hilang (karena di golang tidak wajib menggunakan var)

	namess := "Bima Bagaskhoro"
	fmt.Println(namess)

	fmt.Println(pembatas)

	// Deklarasi Multiple Variable

	var (
		firstName   = "Bima"
		lastName    = "Bagaskhoro"
		declaration = "<---(ini Deklarasi Multiple Variable)"
	)

	fmt.Println(firstName, lastName, declaration)

}
