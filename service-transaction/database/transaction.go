package database

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Transaction struct {
	ID int `gorm:"primary_key" json:"-"`
	Username string `json:"username"`
	MenuName string `json:"menu_name"`
	MenuPrice int `json:"menu_price"`
	MenuQuantity int `json:"menu_quantity"`
	MenuTotalAmount int `json:"menu_total_amount"`
}

func (transaction *Transaction) AddTransaction (db *gorm.DB) error{

	if err := db.Create(transaction).Error;err!=nil {
		return err
	}

	return nil
}

func GetTransactions (user string, db *gorm.DB) ([]Transaction,error){
	var transactions []Transaction

	if err := db.Where(&Transaction{Username: user}).Find(&transactions).Error;err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil,errors.Errorf("user not found")
		}
	}

	return transactions,nil
}
