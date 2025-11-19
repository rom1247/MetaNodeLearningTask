package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Balance float64
}

// addBalance 加余额
func (a *Account) addBalance(amount float64) {
	a.Balance += amount
}

// subBalance 减余额
func (a *Account) subBalance(amount float64) {
	a.Balance -= amount
}
