package main

import (
	"fmt"
	"time"
)

func foo(array [1e6]int) {
	fmt.Println("数组长度: ", len(array))
}

func main() {
	start := time.Now()

	var array [1e6]int
	foo(array)

	elapse := time.Now().Sub(start)
	fmt.Println("耗时(ms)： ", elapse.Nanoseconds()/1000000)

	var slice1 []int
	var slice2 = make([]int, 0)

	slice3 := append(slice2, 2)
	slice4 := append(slice1, 1)

	fmt.Println("slice1 len: ", len(slice1))
	fmt.Println("slice2 len: ", len(slice2))
	fmt.Println("slice3 len: ", len(slice3))
	fmt.Println("slice4 len: ", len(slice4))

	slice := []int{10, 20, 30, 40}
	fmt.Println("slice len:", len(slice), "cap:", cap(slice))

	slice = append(slice, 50)
	fmt.Println("slice len:", len(slice), "cap:", cap(slice))

	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])
	}
}
