package main

import (
	"fmt"
	"log"
	"os"
)

var filename = "test.txt"

func main() {
	writeFile()
}

func ReadFile() {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type : %T | String : %s\n", data, string(data))

}

func Stats() {
	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	/*
		type FileInfo interface {
			Name() string       // base name of the file
			Size() int64        // length in bytes for regular files; system-dependent for others
			Mode() FileMode     // file mode bits
			ModTime() time.Time // modification time
			IsDir() bool        // abbreviation for Mode().IsDir()
			Sys() any           // underlying data source (can return nil)
		}
	*/
	fmt.Println(info.Name(), info.Size(), info.Mode(), info.ModTime(), info.IsDir(), info.Sys())
}

func IsNotExist() {
	_, err := os.Stat("notexist.json")
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	}
}

func create() {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func open() {
	file, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_APPEND, // flags
		0644,                    // permissions
	)
	if err != nil {
		log.Fatal(err)
	}

	// you can use the file (of type os.File) to read, write, Seek etc

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func writeFile() {
	err := os.WriteFile(filename, []byte("New text"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
