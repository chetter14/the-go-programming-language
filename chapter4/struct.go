package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

type Point struct {
	/* Fields are exported */
	X, Y int
}

type address struct {
	hostname string
	port     int
}

func main() {
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior " + *position

	var emplMonth *Employee = &dilbert
	emplMonth.Position += " (proactive team player)"

	fmt.Println(dilbert)

	a := Point{X: 2, Y: 1}
	fmt.Println(a)

	fmt.Println(Scale(Point{1, 2}, 5))

	pp := &Point{4, 3}
	/* Equivalent to:
	pp := new(Point)
	*pp = Point{4, 3}
	*/
	fmt.Println(*pp)

	p := Point{1, 2}
	q := Point{2, 1}
	w := Point{1, 2}
	fmt.Println(q == p)
	fmt.Println(w == p)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
