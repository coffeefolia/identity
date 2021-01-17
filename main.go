package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fname := os.Args[1]
	file, err := os.Open(fname)
	check(err)

}
