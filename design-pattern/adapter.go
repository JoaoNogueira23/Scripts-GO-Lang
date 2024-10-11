package main

import "fmt"

// Interface antiga
type oldPayment interface {
	pay(value float32)
}

// Estrutura do Adapter
type paymentNewAdapter struct {
	newProvider newPayment
}

// Método adaptado para funcionar com o novo provedor (Adapter)
func (a *paymentNewAdapter) PaymenteValue(value float32) {
	a.newProvider.ProcessPayment(value, "BRL") // Agora deve funcionar corretamente
}

// Implementação do novo provedor
type newPayment struct{}

func (n *newPayment) ProcessPayment(total float32, coin string) {
	fmt.Printf("Processing payment of %.2f %s using newPaymentProvider\n", total, coin)
}

func main() {
	// Exemplo de uso
	adapter := &paymentNewAdapter{
		newProvider: newPayment{},
	}

	adapter.PaymenteValue(100.00)
}
