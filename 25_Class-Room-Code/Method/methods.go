package main

import "fmt"

type Direction uint8

const (
	North Direction = iota
	South
	East
	West
)

func (a Direction) String() (ret string) {
	switch a {
	case North:
		ret = "North"
	case South:
		ret = "South"
	case East:
		ret = "East"
	case West:
		ret = "West"
	}
	return
}

func main() {
	fmt.Println("Introduction to Methods")

	var d Direction = South
	var e Direction = East

	fmt.Println(d)
	e.String()

}
