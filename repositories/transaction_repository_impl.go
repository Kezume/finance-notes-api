package repositories

import (
	"github.com/Kezume/finance-notes/database"
	"github.com/Kezume/finance-notes/helpers"
	"github.com/Kezume/finance-notes/models"
)

type TransacationRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransacationRepositoryImpl{}
}

// Create implements TransactionRepository.
func (t *TransacationRepositoryImpl) Create(transaction *models.Transaction) error {
	err := database.DB.Create(transaction).Error

	helpers.IfPanicError(err)

	return err
}

