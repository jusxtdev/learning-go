package main

import (
	"flag"
	"fmt"

	"github.com/jusxtdev/gofio"
)

/*
Word wise analysis of text
- most repeated words
- Longest word
- Total length
*/

func main() {
	// flags
	file := flag.String("file", "", "File to analyze")

	// parse flag
	flag.Parse()

	// err check
	if *file == "" {
		return
	}

	// file handler
	fh := gofio.Gofio{}

	// parse file
	parse_file(&fh, *file)

	// read data
	data := read_file(fh)
	fmt.Println(data)

}

func parse_file(fh *gofio.Gofio, filepath string) {
	err := fh.Initialize(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = fh.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func read_file(fh gofio.Gofio) string {
	content, err := fh.Read()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return content
}
