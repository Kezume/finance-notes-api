package repositories

import "github.com/Kezume/finance-notes/models"

type TransactionRepository interface {
	GetAll() ([]models.TransactionResponse, error)
	Create(transaction *models.Transaction) error
}