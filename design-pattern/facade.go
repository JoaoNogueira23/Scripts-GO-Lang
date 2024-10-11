package main

import (
	"fmt"
)

// PaymentSystem é um subsistema que processa o pagamento.
type PaymentSystem struct{}

// ProcessPayment processa o pagamento de um valor específico.
func (p *PaymentSystem) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Payment of %.2f processed successfully.", amount)
}

// StockSystem é um subsistema que verifica o estoque de produtos.
type StockSystem struct{}

// CheckStock verifica se o produto está disponível em estoque.
func (s *StockSystem) CheckStock(productID string) bool {
	fmt.Printf("Checking stock for product %s...\n", productID)
	return true // Simula que o produto está em estoque
}

// DeliverySystem é um subsistema que organiza a entrega dos produtos.
type DeliverySystem struct{}

// ArrangeDelivery organiza a entrega para o endereço especificado.
func (d *DeliverySystem) ArrangeDelivery(address string) string {
	return fmt.Sprintf("Delivery arranged to address: %s.", address)
}

// OnlineStoreFacade é a Facade que simplifica as interações com os subsistemas.
type OnlineStoreFacade struct {
	paymentSystem  *PaymentSystem
	stockSystem    *StockSystem
	deliverySystem *DeliverySystem
}

// NewOnlineStoreFacade cria uma nova instância da Facade com todos os subsistemas.
func NewOnlineStoreFacade() *OnlineStoreFacade {
	return &OnlineStoreFacade{
		paymentSystem:  &PaymentSystem{},
		stockSystem:    &StockSystem{},
		deliverySystem: &DeliverySystem{},
	}
}

// PlaceOrder coordena o processo de compra verificando o estoque, processando o pagamento e organizando a entrega.
func (store *OnlineStoreFacade) PlaceOrder(productID string, amount float64, address string) string {
	// Verificar estoque
	if !store.stockSystem.CheckStock(productID) {
		return "Product is out of stock."
	}

	// Processar pagamento
	paymentStatus := store.paymentSystem.ProcessPayment(amount)

	// Organizar entrega
	deliveryStatus := store.deliverySystem.ArrangeDelivery(address)

	return fmt.Sprintf("%s\n%s", paymentStatus, deliveryStatus)
}

func main() {
	// Criar a Facade da loja online
	store := NewOnlineStoreFacade()

	// Cliente faz um pedido usando a Facade
	result := store.PlaceOrder("product123", 100.50, "123 Main St.")
	fmt.Println(result)
}
