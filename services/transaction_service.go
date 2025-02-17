package services

import "github.com/Kezume/finance-notes/models"

type TransactionService interface {
	GetAllTransaction() ([]models.TransactionResponse, error)
	CreateTransaction(request models.TransactionRequest) (*models.TransactionResponse, error)
}