package main

import (
	"fmt"
	"github.com/ChrisScotMartin/ait-d/pkgmodd"
)

func main() {
	fmt.Println("Hello World")
	pkgmodd.Hello()
}

func SomeServerFunc() {
	fmt.Println("I'm a func from the main root module ait-d")
}
