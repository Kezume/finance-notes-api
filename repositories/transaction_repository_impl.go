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

// GetAll implements TransactionRepository.
func (t *TransacationRepositoryImpl) GetAll() ([]models.TransactionResponse, error) {
	var transaction []models.TransactionResponse
	err := database.DB.Find(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}


// Create implements TransactionRepository.
func (t *TransacationRepositoryImpl) Create(transaction *models.Transaction) error {
	err := database.DB.Create(transaction).Error

	helpers.IfPanicError(err)

	return err
}
