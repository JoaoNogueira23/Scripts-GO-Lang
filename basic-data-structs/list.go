package main

//imports

import (
	"container/list" // cria listas duplamente encadeadas
	"fmt"
)

func main() {
	var intList list.List

	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)

	for element := intList.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
