package main

import "fmt"

// import "rsc.io/quote"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
