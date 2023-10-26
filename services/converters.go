package services

import (
	"github.com/jt/books/gen/books"
	"github.com/jt/books/models"
	"time"
)

func convertBooksToGoaModel(sc *models.Book) *books.Book {
	if sc == nil || sc.Model.ID == 0 { // this means that the model passed is not correct
		return nil
	}
	id := int(sc.Model.ID)

	var formattedPublishedDate string
	if sc.PublishedAt != nil {
		formattedPublishedDate = sc.PublishedAt.Format(time.RFC3339)
	}
	return &books.Book{
		ID:          &id,
		Title:       &sc.Title,
		Author:      &sc.Author,
		Cover:       EncodeBytesToBase64String(sc.Cover),
		PublishedAt: &formattedPublishedDate,
	}
}

func convertCreateBookPayloadToBookModel(sc *books.CreateBookPayload, coverBytes *[]byte) models.Book {
	if sc == nil {
		return models.Book{}
	}
	var publishedDate time.Time
	if sc.PublishedAt != nil {
		publishedDate, _ = time.Parse(time.RFC3339, *sc.PublishedAt) // since we have checked the format before this we can ignore the error here
	}
	return models.Book{
		Title:       sc.Title,
		Author:      sc.Author,
		Cover:       coverBytes,
		PublishedAt: &publishedDate,
	}
}

func convertUpdatePayloadToBookModel(sc *books.UpdateBookPayload, coverBytes *[]byte) models.Book {
	var publishedDate time.Time
	if sc.PublishedAt != nil {
		publishedDate, _ = time.Parse(time.RFC3339, *sc.PublishedAt) // since we have checked the format before this we can ignore the error here
	}
	model := models.Book{
		Title:       sc.Title,
		Author:      sc.Author,
		Cover:       coverBytes,
		PublishedAt: &publishedDate,
	}
	model.ID = uint(sc.ID)
	return model
}
