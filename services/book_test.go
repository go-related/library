package services

import (
	"github.com/jt/books/models"
	"testing"
)

func TestBookHappyPath(t *testing.T) {

	t.Run("BooksListReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Testing that method List returns correct result")
		var toConvertModel models.Book

		//execute
		result := convertBooksToGoaModel(&toConvertModel)

		//assertion
		if result != nil {
			t.Fatal("failed to return proper response on empty model")
		}

	})
}

func TestBookUnHappyPath(t *testing.T) {

	t.Run("CheckConvertingEmptyBookModelReturnsEmptyResult", func(t *testing.T) {
		//setup
		t.Log("Checking that method convertBooksToGoaModel returns nil on empty source")
		var toConvertModel models.Book

		//execute
		result := convertBooksToGoaModel(&toConvertModel)

		//assertion
		if result != nil {
			t.Fatal("failed to return proper response on empty model")
		}

	})
}
