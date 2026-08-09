package main

import "fmt"

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
