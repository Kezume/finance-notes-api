package services

import (
	"errors"

	"github.com/Kezume/finance-notes/models"
	"github.com/Kezume/finance-notes/repositories"
)

type TransactionServiceImpl struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{repo: repo}
}

// CreateTransaction implements TransactionService.
func (t *TransactionServiceImpl) CreateTransaction(request models.TransactionRequest) (*models.TransactionResponse, error) {
	if request.Amount <= 0 {
		return nil, errors.New("amount harus lebih dari 0")
	}

	transaction := models.Transaction{
		Title: request.Title,
		Amount: request.Amount,
		Type: request.Type,
	}

	err := t.repo.Create(&transaction)
	if err != nil {
		return nil, err
	}

	response := models.TransactionResponse{
		ID: uint(transaction.ID),
		Title: transaction.Title,
		Amount: transaction.Amount,
		Type: transaction.Type,
	}

	return &response, nil
}

