package models

import "errors"

const maxLengthInTitle = 400

//Book struct
type Book struct {
	BookID     int64  `json:"bookID"` //Autoincremental in db2
	Title      string `json:"title"`
	Author     string `json:"author"`
	ReadStatus bool   `json:"readStatus"`
}

//CreateBookCMD struct to create a new book
type CreateBookCMD struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	ReadStatus bool   `json:"readStatus"`
}

func (cmd *CreateBookCMD) validate() error {
	if len(cmd.Title) > maxLengthInTitle {
		return errors.New("Maximum characters exceeded")
	} else if len(cmd.Title) < 1 {
		return errors.New("The title can't be empty")
	}

	return nil
}

//DeleteBookCMD struct to create a new book
type DeleteBookCMD struct {
	BookID int64 `json:"bookId"`
}

//UpdateBookCMD struct to create a new book
type UpdateBookCMD struct {
	BookID     int64  `json:"bookId"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	ReadStatus bool   `json:"readStatus"`
}

func (cmd *UpdateBookCMD) validateUpdate() error {
	if len(cmd.Title) > maxLengthInTitle {
		return errors.New("Maximum characters exceeded")
	} else if len(cmd.Title) < 1 {
		return errors.New("The title can't be empty")
	}

	return nil
}
