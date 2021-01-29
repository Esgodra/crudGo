package gateway

import (
	"crudGo/internal/database"
	"crudGo/models"
)

//BookDeleteGateway interface
type BookDeleteGateway interface {
	Delete(cmd *models.DeleteBookCMD) (*models.Book, error)
}

//BookDeleteGtw struct
type BookDeleteGtw struct {
	BookStorageGateway
}

//NewBookDeleteGateway func
func NewBookDeleteGateway(client *database.DB2Client) BookDeleteGateway {
	return &BookDeleteGtw{&BookStorage{client}}
}
