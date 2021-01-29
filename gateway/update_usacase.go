package gateway

import "crudGo/models"

//Update book method
func (b *BookUpdateGtw) Update(cmd *models.UpdateBookCMD) (*models.Book, error) {
	return b.update(cmd)
}
