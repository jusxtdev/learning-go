## Object Oriented Programming in Golang

### Inheritance
- Parent => `Character`
- Child => `Anime_char` and `Movie_char`
```go
type Character struct {
	name string
	show string
}

type Anime_char struct {
	Character
	studio string
}

type Movie_char struct {
	Character
	actor string
}

func (c Character) info() {
	fmt.Printf("Name : %s | Show : %s\n", c.name, c.show)
}

func (a Anime_char) get_studio() {
	fmt.Printf("Studio : %s\n", a.studio)
}

func (m Movie_char) get_actor() {
	fmt.Printf("Actor : %s\n", m.actor)
}

func main() {
	b := Movie_char{Character{name: "Batman", show: "Dark Knight"}, "Christian Bale"}
	b.info()
	b.get_actor()

	z := Anime_char{Character{name: "Zoro", show: "One Piece"}, "Toei Animation"}
	z.info()
	z.get_studio()

}
```

### Encapsulation
- To encapsulate certain properties of class, put them in another file and *do not export* those properties
- In `./hero/hero.go` -> Declared class, Getters and Setters
- In `./main.go` -> Imported `Hero` class and used it

### Polymorphism & Abstraction
- **Polymorphism** => when two different classes can be treated as the member of the same superclass
    - Here the `Rect` and `Circle` class have the same two methods `area()` and `circumference()`, so we combined that into a single interface called `Shape` where a shape will *should* have the two methods specify in the declaration of interface `Shape`
- **Abstraction** => When the underlying things are abstracted away ie the function doesn't need to worry what it is dealing with
    - The function `print_info` does not need to worry what class does the `s Shape` belong to, it **knows** (from the interface `Shape`) that `s` *will* have the two methods `area()` and `circumference()`

```go
type Shape interface {
	area() float64
	circumference() float64
}

type Rect struct {
	length  int
	breadth int
}

func (r Rect) area() float64 {
	return float64(r.length * r.breadth)
}
func (r Rect) circumference() float64 {
	return float64(r.length + r.breadth)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * (c.radius * c.radius)
}
func (c Circle) circumference() float64 {
	return 2 * 3.14 * c.radius
}

func print_info(s Shape) {
	fmt.Printf("Area : %f | Circumf : %f \n", s.area(), s.circumference())
}

func main() {
	r := Rect{length: 12, breadth: 23}
	print_info(r)

	c := Circle{radius: 32.12}
	print_info(c)
}
```

#### Two more things in Polymorphism
- There are two things supported by Polymorphism
1. Method overriding
    - When child class implements it's own version of a method 
2. Method overloading
    - Passing variable number of arguments to a method / function
