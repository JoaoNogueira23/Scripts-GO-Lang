// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

// main method
func main() {
	fmt.Printf("Hello Workd\n")
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
