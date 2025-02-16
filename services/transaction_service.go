package services

import "github.com/Kezume/finance-notes/models"

type TransactionService interface {
	CreateTransaction(request models.TransactionRequest) (*models.TransactionResponse, error)
}