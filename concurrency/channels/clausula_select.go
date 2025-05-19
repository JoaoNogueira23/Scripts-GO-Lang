package main

import (
	"fmt"
	"math/rand"
)

func sum(total chan int, exit chan bool) {
	value := rand.Intn(20)

	for {
		select {
		case total <- value:
			value += rand.Intn(20)

		case <-exit:
			fmt.Println("Saindo da rotina")
			return
		}
	}
}

func main() {
	total := make(chan int)
	exit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("execução " + fmt.Sprint(i))
			fmt.Println(<-total)
		}
		exit <- true
	}()
	// perceba que ele vai rodar todo o loop do range e depois que fará as tasks definidas na função sum
	// o que na verdade a leitura do que cada channel recebeu durante o loop, mas só ocorre depois em razão da func sum ser chamada depois.
	fmt.Println("saiu do loop na thread principal")
	sum(total, exit)
}
