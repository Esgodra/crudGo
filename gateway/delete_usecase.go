package gateway

import "crudGo/models"

//Delete book method
func (b *BookDeleteGtw) Delete(cmd *models.DeleteBookCMD) (*models.Book, error) {
	return b.delete(cmd)
}
