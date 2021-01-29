package models

import "testing"

func NewBook(title string, author string, readStatus bool) *CreateBookCMD {
	return &CreateBookCMD{
		Title:      title,
		Author:     author,
		ReadStatus: readStatus,
	}
}

func Test_withCorrectParams(t *testing.T) {
	r := NewBook("El Principito", "", true)

	err := r.validate()

	if err != nil {
		t.Error("The validation did not pass")
		t.Fail()
	}
}

func Test_shouldFaildTitle(t *testing.T) {
	r := NewBook("", "", false)

	err := r.validate()

	if err == nil {
		t.Error("Should be fail with empty title")
		t.Fail()
	}
}

func UpdateBook(bookID int64, title string, author string, readStatus bool) *UpdateBookCMD {
	return &UpdateBookCMD{
		BookID:     bookID,
		Title:      title,
		Author:     author,
		ReadStatus: readStatus,
	}
}

func Test_withCorrectParamsUpdate(t *testing.T) {
	r := UpdateBook(5, "El Principito", "", true)

	err := r.validateUpdate()

	if err != nil {
		t.Error("The validation did not pass")
		t.Fail()
	}
}

func Test_shouldFailedTitle(t *testing.T) {
	r := UpdateBook(5, "", "", false)

	err := r.validateUpdate()

	if err == nil {
		t.Error("Should be fail with empty title")
		t.Fail()
	}
}
