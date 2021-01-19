package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return "ee" + string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 200; i++ {
		fmt.Print(randSeq(3) + "\t")
	}
}
