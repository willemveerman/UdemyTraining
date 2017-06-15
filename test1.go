package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}

func b() {
	x := 5
	zero(x)
	fmt.Errorf(x) // x is still 5
}
