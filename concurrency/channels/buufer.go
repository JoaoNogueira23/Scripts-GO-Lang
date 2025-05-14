package main

import (
	"fmt"
	"time"
)

func numbers(ch chan int) {
	for i := 1; i < 11; i++ {
		ch <- i
		fmt.Printf("valor %d enviado para o channel\n", i)
	}

	close(ch)
}

func main() {
	// o segundo parâmetro sedine o buffer do canal, ou seja,
	// quantos valores serão enviados para o canal antes de bloquear
	// bloqueia o envio de valores quando atingi a capacidade, nessa caso 3
	// desbloqueia "espaço" quando um valor é lido do canal
	// Funciona como uma emergência hospitalar, bloqueia o acesso a atendimento quando
	// todos os médicos estão ocupados, mas libera o acesso quando
	// um médico termina
	ch := make(chan int, 3)

	go numbers(ch)

	time.Sleep(200 * time.Millisecond)

	for value := range ch {
		fmt.Printf("valor %d lido do channel\n", value)
		time.Sleep(200 * time.Millisecond)
	}

}
