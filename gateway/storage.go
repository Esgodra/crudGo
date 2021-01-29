package gateway

import (
	"crudGo/internal/database"
	"crudGo/models"
	"fmt"
)

//BookStorageGateway interface
type BookStorageGateway interface {
	create(cmd *models.CreateBookCMD) (*models.Book, error)
	update(cmd *models.UpdateBookCMD) (*models.Book, error)
	delete(cmd *models.DeleteBookCMD) (*models.Book, error)
	get() ([]*models.Book, error)
}

//BookStorage struct
type BookStorage struct {
	*database.DB2Client
}

func (b *BookStorage) create(cmd *models.CreateBookCMD) (*models.Book, error) {
	tx, err := b.DB2Client.Begin()

	if err != nil {
		fmt.Printf("Cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into books (title,author,readstatus) values(?,?,?)`, cmd.Title, cmd.Author, cmd.ReadStatus)

	if err != nil {
		fmt.Printf("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	rowsAff, err := res.RowsAffected()

	if rowsAff != 1 {
		fmt.Printf("cannot fetch last id %s", err)
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Book{
		Title:      cmd.Title,
		Author:     cmd.Author,
		ReadStatus: cmd.ReadStatus,
	}, nil
}

func (b *BookStorage) update(cmd *models.UpdateBookCMD) (*models.Book, error) {
	tx, err := b.DB2Client.Begin()

	if err != nil {
		fmt.Printf("Cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`update books set title=?, author=?,readstatus=? where bookid=?`, cmd.Title, cmd.Author, cmd.ReadStatus, cmd.BookID)

	if err != nil {
		fmt.Printf("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	rowsAff, err := res.RowsAffected()

	if rowsAff != 1 {
		fmt.Printf("cannot update this book")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Book{
		BookID:     cmd.BookID,
		Title:      cmd.Title,
		Author:     cmd.Author,
		ReadStatus: cmd.ReadStatus,
	}, nil
}

func (b *BookStorage) delete(cmd *models.DeleteBookCMD) (*models.Book, error) {
	tx, err := b.DB2Client.Begin()

	if err != nil {
		fmt.Printf("Cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`delete from books where bookid=?`, cmd.BookID)

	if err != nil {
		fmt.Printf("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	rowsAff, err := res.RowsAffected()

	if rowsAff != 1 {
		fmt.Printf("cannot delete this book")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Book{
		BookID: cmd.BookID,
	}, nil
}

func (b *BookStorage) get() ([]*models.Book, error) {
	var books []*models.Book

	res, err := b.Query(`select * from books`)

	if err != nil {
		fmt.Printf("cannot execute statement %s", err)
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var b *models.Book
		err = res.Scan(&b.BookID, &b.Title, &b.Author, &b.ReadStatus)
		if err != nil {
			fmt.Printf("Scan: %s", err)
		}
		books = append(books, b)
	}
	fmt.Println(books)
	return books, nil
}
