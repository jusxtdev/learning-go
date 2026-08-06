# Golang Notes

### `flags`
```go
```go 
name = flag.String("name", "Guest", "Name of the User")

flag.Parse()

fmt.Printf("Hello, %s\n", *name)

// execute as ./main -name Dev 
// Hello, Dev
```


### `os`
- Refer `./os-package/`
- Some common patterns

1. Create a file if it doesn't exist and initialize it  
```go 
func createFile(filename string){
	// If the file doesn't exist, create it, or append to the file
	var file *os.File
	var err error

	// return if file already exists
	if _, err = os.Stat(filename); !os.IsNotExist(err) {
		return
	}

	// create new file and
	file, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	// initialize it with '[]'
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
```

### `json`
- **High Level Process for working with json data**
    1. Read json data into the memory (`os.ReadFile`)
    2. Un-marshal the json data into structs and slices
    3. Work with that (append, delete ...)
    4. Marshal the data into json
    5. Write the json data back to the file
> Make sure the struct fields prepared for containing json data
1. Make the strcut fields exported by capitalizing them 
2. add proper `json:"field_name"` tags 
> OTHERWISE : The `Marshal` / `MarshalIndent` will return an empty struct (So fuc*ing important)
- Example
```go
type Name struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

// variable to store the json file data
var people []Name
```
