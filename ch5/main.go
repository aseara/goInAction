package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func main() {
	fmt.Println("Hello Go!")

	var bill user
	lisa := user{
		name:       "lisa",
		email:      "lisa@163.com",
		ext:        123,
		privileged: true,
	}

	fmt.Println("bill", bill)
	fmt.Println("lisa", lisa)

}
