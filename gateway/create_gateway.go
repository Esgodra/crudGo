package gateway

import (
	"crudGo/internal/database"
	"crudGo/models"
)

//BookCreateGateway interface
type BookCreateGateway interface {
	Create(cmd *models.CreateBookCMD) (*models.Book, error)
}

//BookCreateGtw struct
type BookCreateGtw struct {
	BookStorageGateway
}

//NewBookCreateGateway func
func NewBookCreateGateway(client *database.DB2Client) BookCreateGateway {
	return &BookCreateGtw{&BookStorage{client}}
}
