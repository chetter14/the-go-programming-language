package main

import (
	"fmt"
	"regexp"
)

func main() {
	src := "ls; echo $bar; $foo -v; x = $bar; $exec123 -D > file.txt; $ping 192.168.0.1"

	f := func(s string) string {
		if s == "foo" {
			return "FOO_CMD"
		} else if s == "bar" {
			return "GLOBAL_CONFIG_BAR"
		} else if s == "exec123" {
			return "/home/user/exec123"
		} else {
			return "UNDEFINED_STRING"
		}
	}

	fmt.Println(src)
	fmt.Println(expand(src, f))
}

func expand(s string, f func(string) string) string {
	varPattern := regexp.MustCompile(`\$[A-Za-z_][A-Za-z_0-9]*`)

	return varPattern.ReplaceAllStringFunc(s, func(match string) string {
		return f(match[1:])
	})
}
