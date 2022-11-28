package main

import "fmt"

const englishhelloprefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishhelloprefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
