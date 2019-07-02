package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("note.md", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var i io.Reader = f
	switch v := i.(type) {
	case *os.File, io.ReadWriter:
		if v == i {
			fmt.Println(true)
		}
	default:
		return
	}
}
