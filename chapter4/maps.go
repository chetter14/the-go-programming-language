package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)

	ages = map[string]int{
		"bob":   24,
		"alice": 31,
	}
	ages["chris"] = 15

	fmt.Println(ages)
	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages)
	fmt.Println(ages["alice"])

	ages["alice"] = ages["alice"] + 1
	ages["bob"]++
	ages["abby"] = 42
	fmt.Println(ages)

	// a := &ages["alice"]		// invalid operation! can't take address

	for name, age := range ages {
		fmt.Printf("%s %d\n", name, age)
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	/* map stores 'nil' value */
	var a map[string]int
	fmt.Println(a == nil)
	fmt.Println(len(a) == 0)

	/* map is just empty here */
	var b = map[string]int{}
	fmt.Println(b != nil)
	fmt.Println(len(b) == 0)

	if val, ok := ages["john"]; !ok {
		fmt.Printf("John name is not present in map! val=%d\n", val)
	}

	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
	
	list := []string{"hello", "world", "!!!"}
	str := fmt.Sprintf("%q", list)
	fmt.Println(str)
	fmt.Println(list)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
