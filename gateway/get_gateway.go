package gateway

import (
	"crudGo/internal/database"
	"crudGo/models"
	"fmt"
)

//BookGetGateway interface
type BookGetGateway interface {
	Get() ([]*models.Book, error)
}

//BookGetGtw struct
type BookGetGtw struct {
	BookStorageGateway
}

//NewBookGetGateway func
func NewBookGetGateway(client *database.DB2Client) BookGetGateway {
	fmt.Println("Gateway")
	return &BookGetGtw{&BookStorage{client}}
}
