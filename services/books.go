package services

import (
	"context"
	"errors"
	"github.com/jt/books/database"
	"github.com/jt/books/gen/books"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// books service example implementation.
// The example methods log the requests and return zero values.
type bookssrvc struct {
	logger *logrus.Logger
	db     database.BooksDB
}

// NewBooks returns the books service implementation.
func NewBooks(lg *logrus.Logger, db database.BooksDB) books.Service {
	return &bookssrvc{logger: lg, db: db}
}

func (s *bookssrvc) List(ctx context.Context) (res []*books.Book, err error) {
	// do the request on db
	result, err := s.db.GetAllBooks(ctx)
	if err != nil {
		s.logger.WithError(err).Error("couldn't load books")
		return
	}
	//prepare the output
	for _, book := range result {
		res = append(res, convertBooksToGoaModel(book))
	}
	return
}

func (s *bookssrvc) Show(ctx context.Context, payload *books.ShowPayload) (res *books.Book, err error) {
	// request data on db
	result, err := s.db.GetBookById(ctx, uint(payload.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //special case we don't need to log this
			return nil, books.MakeNotFound(err)
		}
		s.logger.WithError(err).Error("couldn't load book")
		return
	}
	// prepare the response
	res = convertBooksToGoaModel(result)
	return
}

// Create a new book
func (s *bookssrvc) Create(ctx context.Context, p *books.CreateBookPayload) (res *books.Book, err error) {
	imageBytes, err := DecodeBase64StringToBytes(p.Cover)
	if err != nil {
		s.logger.WithError(err).Error("couldn't decode cover to b64")
		return nil, books.MakeBadRequest(err)
	}

	toUpdateModel := convertCreateBookPayloadToBookModel(p, imageBytes)
	result, err := s.db.CreateBook(ctx, toUpdateModel)
	if err != nil {
		s.logger.WithError(err).Error("couldn't create book")
		return
	}
	res = convertBooksToGoaModel(&result)
	return
}

// Update an existing book
func (s *bookssrvc) Update(ctx context.Context, p *books.UpdateBookPayload) (res *books.Book, err error) {
	imageBytes, err := DecodeBase64StringToBytes(p.Cover)
	if err != nil {
		s.logger.WithError(err).WithField("id", p.ID).Error("couldn't decode cover to b64")
		return nil, books.MakeBadRequest(err)
	}

	toUpdateModel := convertUpdatePayloadToBookModel(p, imageBytes)
	err = s.db.UpdateBook(ctx, toUpdateModel)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, books.MakeNotFound(err)
		}
		s.logger.WithError(err).WithField("id", p.ID).Error("couldn't update book")
		return
	}
	res = convertBooksToGoaModel(&toUpdateModel)
	return
}

// Delete a book by ID
func (s *bookssrvc) Delete(ctx context.Context, p *books.DeletePayload) (err error) {
	err = s.db.DeleteBook(ctx, uint(p.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return books.MakeNotFound(err) //special error
		}
		s.logger.WithError(err).WithField("id", p.ID).Error("couldn't delete book")
	}

	return
}
