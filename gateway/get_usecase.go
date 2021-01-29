package gateway

import "crudGo/models"

//Get book method
func (b *BookGetGtw) Get() ([]*models.Book, error) {
	return b.get()
}
