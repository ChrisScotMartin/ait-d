package pkgmodd

import "fmt"

func Hello() {
	fmt.Println("From pkgmodd")
	fmt.Println("Hello, I'm either a package or a module...depends on version.")
}
