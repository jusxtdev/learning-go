package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

/*
1. Read json data into the memory (`os.ReadFile`)
2. Un-marshal the json data into structs and slices
3. Work with that (append, delete ...)
4. Marshal the data into json
5. Write the json data back to the file
*/

// what each entry looks like
// NOTE - tags are sooo important
// 1. Make sure the struct fields are exported
// 2. add tags
// IMPORTANT -> or else Marshal / MarshalIndent will return empty structs
type Name struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

// variable to store the json file data
var people []Name

// json file
var filename string = "file.json"

func main() {
	create_json_file(filename)

	read_json_file(filename)

	name := flag.String("name", "", "Name of the person")
	age := flag.Float64("age", 0.0, "Age of the person")
	flag.Parse()
	if *name != "" || *age != 0.0 {
		add_data(*name, *age)
	}

	fmt.Println("IN main : ", people)

	update_json_file(filename)
}

func create_json_file(filename string) {
	var file *os.File
	var err error

	if _, err = os.Stat(filename); !os.IsNotExist(err) {
		// file already exists
		return
	}

	file, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte("[]"))
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

func read_json_file(filename string) {
	// read json data
	// data => []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal and store json data as a slice or struct
	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Fatal(err)
	}
}

func add_data(name string, age float64) {
	new_person := Name{
		Name: name,
		Age:  age,
	}
	people = append(people, new_person)
	fmt.Println("IN add_data : ", people)
}

func update_json_file(filename string) {
	fmt.Println("IN update before marshal : ", people)
	data, err := json.MarshalIndent(people, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IN update after marshal : ", string(data))

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
