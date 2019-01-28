package main

import "fmt"

type duration int

func (d duration) pretty() {
	fmt.Printf("Duration Address: %d", &d)
}

func main() {
	duration(42).pretty()
}
