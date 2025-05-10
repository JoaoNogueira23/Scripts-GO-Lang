package main

import (
	"fmt"
)

func numbers(ch chan int) {
	for i := 1; i < 11; i++ {
		ch <- i
	}

	close(ch)
}

func main() {
	// Cria um canal
	ch := make(chan int)

	// iniciando uma goroutine para enviar valores para o canal os números de 1 a 10
	go numbers(ch)

	// loop para receber os valores do canal com condição de parada para o fim da execução do canal
	for {
		value, ok := <-ch
		if !ok {
			break
		}

		fmt.Println(value)
	}
}
