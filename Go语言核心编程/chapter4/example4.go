package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("test.md", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var i io.Reader = f

	switch v := i.(type) {

	case *os.File:
		v.Write([]byte("*os.File\n"))
		v.Sync()
	case io.ReadWriter:
		v.Write([]byte("io.ReadWriter\n"))

	default:
		return
	}
}
