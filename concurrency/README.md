
# 📘 Goroutines em Go: Execução Concorrente Simples

Este projeto é uma introdução prática ao uso de **goroutines** em Go, que permite a execução concorrente de funções. O código mostra como duas funções (`first_sequence` e `second_sequence`) podem ser executadas em paralelo utilizando o comando `go`.

## 🧠 Conceitos Envolvidos

### ✅ Goroutines

Goroutines são funções ou métodos que rodam **concorrentemente** com outras funções. Elas são extremamente leves, iniciadas com a palavra-chave `go` antes da chamada da função:

---
```go
go nomeDaFuncao()
```
---

Ao contrário de `threads` tradicionais, centenas ou milhares de goroutines podem ser executadas com baixo consumo de recursos.

## 📄 Estrutura do Código

### Função `first_sequence`

---
```go
func first_sequence() {
	numbers := []int{1, 3, 5}
	for _, value := range numbers {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("first routine finished")
}
```
---

* Imprime os números ímpares `1`, `3`, `5` com um intervalo de 200ms entre cada impressão.
* Após finalizar o loop, exibe a mensagem `"first routine finished"`.

### Função `second_sequence`

---
```go
func second_sequence() {
	numbers := []int{2, 4}
	for _, value := range numbers {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(value)
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("Second routine finished")
}
```
---

* Imprime os números pares `2` e `4`.
* Para cada número, espera 100ms antes de imprimir e depois 200ms antes de seguir.

### Função `main`

---
```go
func main() {
	go first_sequence()
	go second_sequence()
	time.Sleep(time.Millisecond * 1101)
	println("Execution finished")
}
```
---

* Ambas as funções são iniciadas como **goroutines**.
* A `main` espera por **1.101ms (1.1 segundos)** para que as goroutines tenham tempo de completar.
* Em seguida, imprime `"Execution finished"`.

️ **Importante**: Se o `Sleep` for removido ou for menor que o tempo de execução das goroutines, o programa pode encerrar antes delas terminarem — porque a função `main` também é uma goroutine especial, e quando ela termina, o programa inteiro encerra.

## ⏱️ Linha do Tempo (aproximada)

| Tempo  | Acontecimento                     |
| ------ | --------------------------------- |
| 0ms    | Ambas as goroutines são iniciadas |
| 100ms  | `2` é impresso                    |
| 200ms  | `1` é impresso                    |
| 300ms  | `3` é impresso                    |
| 400ms  | `4` é impresso                    |
| 500ms  | `5` é impresso                    |
| 700ms  | `first routine finished`          |
| 800ms  | `Second routine finished`         |
| 1101ms | `Execution finished`              |

## ✅ Aprendizados

* Como iniciar **goroutines** com `go`.
* Como o uso de `Sleep` simula operações demoradas e ajuda a entender concorrência.
* A importância de sincronização: sem mecanismos como `WaitGroup`, o tempo de `Sleep` em `main` precisa ser suficiente.

## 🚀 Próximos Passos

Para melhorar esse código e ter controle real da execução, você pode:

* Usar **`sync.WaitGroup`** para aguardar as goroutines terminarem de forma segura.
* Criar **channels** para comunicação entre as rotinas.

## 📊 channels (`chan`) em Go

channels são estruturas que permitem a comunicação segura entre goroutines, funcina como um tunel por onde passa informações entre as goroutines. A sintaxe `chan int` indica um canal que transporta valores do tipo `int`.

### Exemplo com canal:

---
```go
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
```
---

### Explicação:

* `make(chan int)` cria um canal de inteiros.
* A goroutine `numbers` envia os números de 1 a 10 usando `ch <- i`.
* O canal é fechado com `close(ch)` para indicar que não há mais dados.
* No `main`, usamos `value, ok := <-ch` para receber valores até o canal ser fechado.

### Benefícios:

* Garante sincronização entre as goroutines.
* Evita condições de corrida e uso indevido de `sleep` para controlar execução.

> Com canais, o código Go alcança uma forma elegante e segura de concorrência.
