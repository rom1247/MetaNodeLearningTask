package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

//转账

func Transfer(from, to *Account, amount float64) (*Transaction, error) {

	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}
	if from == nil {
		return nil, fmt.Errorf("from account is nil")
	}
	if to == nil {
		return nil, fmt.Errorf("to account is nil")
	}

	if from.Balance < amount {
		return nil, fmt.Errorf("insufficient balance")
	}

	from.subBalance(amount)
	to.addBalance(amount)

	transaction := &Transaction{
		FromAccountId: from.ID,
		ToAccountId:   to.ID,
		Amount:        amount,
	}

	return transaction, nil
}
