package main

import (
	"fmt"
	"strconv"
)

/*
func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
}

func main() {
	chs := make([]chan int, 10)
	for i :=0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		fmt.Println(<- ch)
	}
}*/

func Count(ch chan string, i int) {
	fmt.Println("Counting")
	ch <- strconv.Itoa(i)
}

func main() {
	chs := make(chan string)
	for i := 0; i < 10; i++ {
		go Count(chs, i)
	}
	for j := 0; j < 10; j++ {
		fmt.Println(<-chs)
	}
}
