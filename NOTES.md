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

### `json`

