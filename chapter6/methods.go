package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	a := Point{1, 2}
	b := Point{4, 6}
	fmt.Println(Distance(a, b))
	fmt.Println(a.Distance(b))
	
	distanceFromA := a.Distance		// method value (binds to a specific receiver)
	fmt.Println(distanceFromA(b))
	var origin Point
	fmt.Println(distanceFromA(origin))
	
	dist := Point.Distance			// method expression (supply a receiver as well)
	fmt.Println(dist(a, b))
	fmt.Printf("%T\n", dist)
}

type Point struct{ X, Y float64 }

func Distance(a, b Point) float64 {
	return math.Hypot(b.X-a.X, b.Y-a.Y)
}

func (p Point) Distance(a Point) float64 {
	return math.Hypot(a.X-p.X, a.Y-p.Y)
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(k string) string {
	cache.Lock()
	v := cache.mapping[k]
	cache.Unlock()
	return v
}
