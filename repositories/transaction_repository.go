package repositories

import "github.com/Kezume/finance-notes/models"

type TransactionRepository interface {
	GetAll() ([]models.Transaction, error)
	Create(transaction *models.Transaction) error
}