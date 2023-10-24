package library

import (
	"context"
	"log"

	books "github.com/jt/books/gen/books"
)

// books service example implementation.
// The example methods log the requests and return zero values.
type bookssrvc struct {
	logger *log.Logger
}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger) books.Service {
	return &bookssrvc{logger}
}

// Retrieve all books
func (s *bookssrvc) Items(ctx context.Context) (res []*books.Book, err error) {
	s.logger.Print("books.items")
	return
}

// Retrieve a book by ID
func (s *bookssrvc) Item(ctx context.Context, p *books.ItemPayload) (res *books.Book, err error) {
	res = &books.Book{}
	s.logger.Print("books.item")
	return
}

// Create a new book
func (s *bookssrvc) Create(ctx context.Context, p *books.CreateBookPayload) (res *books.Book, err error) {
	res = &books.Book{}
	s.logger.Print("books.create")
	return
}

// Update an existing book
func (s *bookssrvc) Update(ctx context.Context, p *books.UpdateBookPayload) (res *books.Book, err error) {
	res = &books.Book{}
	s.logger.Print("books.update")
	return
}

// Delete a book by ID
func (s *bookssrvc) Delete(ctx context.Context, p *books.DeletePayload) (err error) {
	s.logger.Print("books.delete")
	return
}
