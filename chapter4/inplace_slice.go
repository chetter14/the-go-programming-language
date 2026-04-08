package main

import "fmt"

func main() {
	var s = []string{"hey", "", "world", "", "!!!"}
	b := nonempty(s)
	fmt.Println(s)
	fmt.Println(b)

	var stack []int
	stack = append(stack, 1, 2, 3)
	fmt.Println(stack)
	top := stack[len(stack)-1]
	fmt.Println(top)
	/* Pop operation */
	stack = stack[:len(stack)-1]
	fmt.Println(stack)

	arr := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(arr, 2))
}

/* Removes empty strings in-place */
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

/* Preserve order */
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
