package main

import "fmt"

func Greeting(name string) string {
	if name == "" {
		return "Hello, world"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(Greeting("foo"))
}
