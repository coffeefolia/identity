package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
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
	f1, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	l1 := strings.Split(string(f1), "\n")
	f2, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	l2 := strings.Split(string(f2), "\n")

	for i := 0; i < 50; i++ {
		a := rand.Intn(len(l1) - 1)
		b := rand.Intn(len(l2) - 1)
		c := l1[a] + l2[b]
		if len(c) > 12 {
			println(c[:12])
		} else {
			println(c)
		}
	}
}
