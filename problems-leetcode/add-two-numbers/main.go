package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addValue(value int, header **ListNode) {
	aux := &ListNode{Val: value}
	if *header == nil {
		*header = aux
	} else {
		current := *header
		for current.Next != nil {
			current = current.Next
		}
		current.Next = aux
	}
}

func printList(header *ListNode) {
	current := header
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func reversedNumberBig(stringNumber string) *big.Int {
	// Reverter a string
	reversed := ""
	length := len(stringNumber)

	for i := length - 1; i >= 0; i-- {
		reversed += string(stringNumber[i]) // Adiciona os caracteres na ordem inversa
	} /* O(n) -> the size of list */

	// Converter a string invertida para *big.Int
	val := new(big.Int)
	val, ok := val.SetString(reversed, 10) // Base 10
	if !ok {
		fmt.Println("Erro ao converter a string para big.Int")
		return nil
	}

	return val
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/* get sum of numbers from list */
	var string1 string
	var string2 string
	for l1 != nil || l2 != nil {
		if l1 != nil {
			string1 += strconv.Itoa(l1.Val)
			l1 = l1.Next
		}

		if l2 != nil {
			string2 += strconv.Itoa(l2.Val)
			l2 = l2.Next
		}
	} /* O(n) -> n is the largest size among the lists*/
	/* GET THE SUM */

	sum_result := new(big.Int).Add(reversedNumberBig(string1), reversedNumberBig(string2))

	string_result := sum_result.String()

	var header *ListNode
	var current *ListNode
	lenght_string := len(string_result)
	for i := 0; i < lenght_string; i++ {
		val, _ := strconv.Atoi(string(string_result[lenght_string-1-i]))
		aux := &ListNode{Val: val}
		if header == nil {
			header = aux
			current = header
		} else {
			current.Next = aux
			current = aux
		}

	} /* O(n) -> n is the largest size among the lists with one or two plus interactions */
	return header
}

func main() {
	var l1 *ListNode
	var l2 *ListNode

	// Construção de l1
	addValue(1, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(0, &l1)
	addValue(1, &l1)

	// Construção de l2
	addValue(5, &l2)
	addValue(6, &l2)
	addValue(4, &l2)

	headerBuild := addTwoNumbers(l1, l2)

	printList(headerBuild)
}
