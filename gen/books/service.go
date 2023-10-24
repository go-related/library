// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books service
//
// Command:
// $ goa gen github.com/jt/books/design

package books

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The book service performs CRUD operations on the books resource.
type Service interface {
	// Retrieve all books
	Items(context.Context) (res []*Book, err error)
	// Retrieve a book by ID
	Item(context.Context, *ItemPayload) (res *Book, err error)
	// Create a new book
	Create(context.Context, *CreateBookPayload) (res *Book, err error)
	// Update an existing book
	Update(context.Context, *UpdateBookPayload) (res *Book, err error)
	// Delete a book by ID
	Delete(context.Context, *DeletePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "books"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"items", "item", "create", "update", "delete"}

// Book is the result type of the books service item method.
type Book struct {
	// ID of the book
	ID *int
	// Title of the book
	Title *string
	// Author of the book
	Author *string
	// Book cover image URL
	Cover *string
	// Published date of the book
	PublishedAt *string
}

// CreateBookPayload is the payload type of the books service create method.
type CreateBookPayload struct {
	// Title of the book
	Title string
	// Author of the book
	Author string
	// Book cover image URL
	Cover *string
	// Published date of the book
	PublishedAt *string
}

// DeletePayload is the payload type of the books service delete method.
type DeletePayload struct {
	// ID of the book
	ID int
}

// ItemPayload is the payload type of the books service item method.
type ItemPayload struct {
	// ID of the book
	ID int
}

// UpdateBookPayload is the payload type of the books service update method.
type UpdateBookPayload struct {
	// ID of the book
	ID int
	// Title of the book
	Title string
	// Author of the book
	Author string
	// Book cover image URL
	Cover *string
	// Published date of the book
	PublishedAt *string
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_found", false, false, false)
}
