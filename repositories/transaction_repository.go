package repositories

import "github.com/Kezume/finance-notes/models"

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
}