package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// seed random with time
	rand.Seed(time.Now().UTC().UnixNano())
	// read in dictionary file
	f, err := ioutil.ReadFile("files/5ltrwd.txt")
	check(err)
	l := strings.Split(string(f), "\n")

	for i := 0; i < 50; i++ { // combine and print word mash
		lenl := len(l) - 1
		a := rand.Intn(lenl - 1)
		b := rand.Intn(lenl - 1)
		c := l[a] + l[b]
		if len(c) > 12 {
			println(c[:12])
		} else {
			println(c)
		}
	}
}
