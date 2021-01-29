package gateway

import "crudGo/models"

//Create book method
func (b *BookCreateGtw) Create(cmd *models.CreateBookCMD) (*models.Book, error) {
	return b.create(cmd)
}
