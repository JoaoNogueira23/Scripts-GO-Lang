package main

import (
	"fmt"
	"time"
)

func first_sequence() {

	numbers := []int{1, 3, 5}
	for _, value := range numbers {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("first routine finished")
}

func second_sequence() {
	numbers := []int{2, 4}
	for _, value := range numbers {
		time.Sleep(time.Millisecond * 50)
		fmt.Println(value)
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("Second routine finished")
}

func main() {

	// inputs user about length, width and height
	go first_sequence()
	go second_sequence()
	time.Sleep(time.Millisecond * 1101)
	println("Execution finished")
}
