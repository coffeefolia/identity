package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const DICT_LOC = "dict/dict.txt"
const WORDLEN_A = 3
const WORDLEN_B = 4

func main() {
	// seed random with time
	rand.Seed(time.Now().UTC().UnixNano())
	f1, err := ioutil.ReadFile(DICT_LOC)
	if err != nil {
		log.Fatal(err)
	}
	l1 := strings.Split(string(f1), "\n")

	for i := 0; i < 500; i++ {
		a := rand.Intn(len(l1) - 1)
		b := rand.Intn(len(l1) - 1)
		c := l1[a]
		d := l1[b]
		if len(c) > WORDLEN_A {
			c = c[:WORDLEN_A]
		}
		if len(c) > WORDLEN_B {
			d = d[:WORDLEN_B]
		}
		e := c + d
		print(e, "\t")

	}
}
