package gateway

import (
	"crudGo/internal/database"
	"crudGo/models"
)

//BookUpdateGateway interface
type BookUpdateGateway interface {
	Update(cmd *models.UpdateBookCMD) (*models.Book, error)
}

//BookUpdateGtw struct
type BookUpdateGtw struct {
	BookStorageGateway
}

//NewBookUpdateGateway func
func NewBookUpdateGateway(client *database.DB2Client) BookUpdateGateway {
	return &BookUpdateGtw{&BookStorage{client}}
}
