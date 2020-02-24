package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message")

	log.Panicln("panic message")

	fmt.Println("end of main")

	log.Fatalln("fatal message")
}
