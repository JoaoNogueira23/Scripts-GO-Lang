package main

import (
	"encoding/json"
	"fmt"
)

// os detalhes da conta fica de uma forma isolada, isso chamamos de private class
type AccountDetails struct {
	id          string
	accountType string
}

type Account struct {
	details      *AccountDetails
	CustomerName string
}

func (account *Account) setDetails(id, accountType string) {
	if account.details == nil {
		account.details = &AccountDetails{} // Inicializa a struct AccountDetails
	}
	account.details = &AccountDetails{id, accountType}
}

func (account *Account) getAccountType() string {
	return account.details.accountType
}

func main() {
	var account *Account = &Account{CustomerName: "João"}
	// Definindo os detalhes da conta
	account.setDetails("12345", "Savings")

	jsonAccount, _ := json.Marshal(account)

	fmt.Println("Private class hidden", string(jsonAccount))

	fmt.Println("Account type: ", account.getAccountType())
}
