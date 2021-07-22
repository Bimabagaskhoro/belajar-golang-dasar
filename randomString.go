package main

import (
	"fmt"
	"math/rand"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {
	var (
		kutip = `"`
		koma  = `,`
	)
	for i := 0; i < 200; i++ {
		fmt.Println(kutip + RandomString(8) + kutip + koma)
	}

}
