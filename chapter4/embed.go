package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	/* anonymous field; embedded within the Circle */
	/* this field internally has a name - that of the named type */
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	fmt.Println("Hey there")

	var w Wheel
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20

	/* error ! */
	// wh := Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20}

	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)
}
