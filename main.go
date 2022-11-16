package main

import (
	"fmt"
	"github.com/ChrisScotMartin/ait-a"
	"github.com/ChrisScotMartin/ait-c/pkgmodc"
)

func main() {

	fmt.Println("Hello World")
	pkgmodd.Hello()
}

func SomeServerFunc() {
	fmt.Println("I'm a func from the main root module ait-d")
}
