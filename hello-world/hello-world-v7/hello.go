package main

import "fmt"

const bengali = "Bengali"
const bengaliHelloPrefix = "হ্যালো, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	if language == bengali {
		return bengaliHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("এহসান!", "Bengali"))
}
