// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books HTTP server types
//
// Command:
// $ goa gen github.com/jt/books/design

package server

import (
	books "github.com/jt/books/gen/books"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "books" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Base64 of the Book cover image
	Cover *string `form:"cover,omitempty" json:"cover,omitempty" xml:"cover,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// UpdateRequestBody is the type of the "books" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Base64 of the Book cover image
	Cover *string `form:"cover,omitempty" json:"cover,omitempty" xml:"cover,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// ListResponseBody is the type of the "books" service "list" endpoint HTTP
// response body.
type ListResponseBody []*BookResponse

// ShowResponseBody is the type of the "books" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID of the book
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Base64 of the Book cover image
	Cover *string `form:"cover,omitempty" json:"cover,omitempty" xml:"cover,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// CreateResponseBody is the type of the "books" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// ID of the book
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Base64 of the Book cover image
	Cover *string `form:"cover,omitempty" json:"cover,omitempty" xml:"cover,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// UpdateResponseBody is the type of the "books" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// ID of the book
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Base64 of the Book cover image
	Cover *string `form:"cover,omitempty" json:"cover,omitempty" xml:"cover,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "books" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateBadRequestResponseBody is the type of the "books" service "create"
// endpoint HTTP response body for the "bad_request" error.
type CreateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateNotFoundResponseBody is the type of the "books" service "update"
// endpoint HTTP response body for the "not_found" error.
type UpdateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateBadRequestResponseBody is the type of the "books" service "update"
// endpoint HTTP response body for the "bad_request" error.
type UpdateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteNotFoundResponseBody is the type of the "books" service "delete"
// endpoint HTTP response body for the "not_found" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// BookResponse is used to define fields on response body types.
type BookResponse struct {
	// ID of the book
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Base64 of the Book cover image
	Cover *string `form:"cover,omitempty" json:"cover,omitempty" xml:"cover,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "books" service.
func NewListResponseBody(res []*books.Book) ListResponseBody {
	body := make([]*BookResponse, len(res))
	for i, val := range res {
		body[i] = marshalBooksBookToBookResponse(val)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "books" service.
func NewShowResponseBody(res *books.Book) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          res.ID,
		Title:       res.Title,
		Author:      res.Author,
		Cover:       res.Cover,
		PublishedAt: res.PublishedAt,
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "books" service.
func NewCreateResponseBody(res *books.Book) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:          res.ID,
		Title:       res.Title,
		Author:      res.Author,
		Cover:       res.Cover,
		PublishedAt: res.PublishedAt,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "books" service.
func NewUpdateResponseBody(res *books.Book) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:          res.ID,
		Title:       res.Title,
		Author:      res.Author,
		Cover:       res.Cover,
		PublishedAt: res.PublishedAt,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "books" service.
func NewShowNotFoundResponseBody(res *goa.ServiceError) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateBadRequestResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "books" service.
func NewCreateBadRequestResponseBody(res *goa.ServiceError) *CreateBadRequestResponseBody {
	body := &CreateBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "books" service.
func NewUpdateNotFoundResponseBody(res *goa.ServiceError) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateBadRequestResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "books" service.
func NewUpdateBadRequestResponseBody(res *goa.ServiceError) *UpdateBadRequestResponseBody {
	body := &UpdateBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "books" service.
func NewDeleteNotFoundResponseBody(res *goa.ServiceError) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowPayload builds a books service show endpoint payload.
func NewShowPayload(id int) *books.ShowPayload {
	v := &books.ShowPayload{}
	v.ID = id

	return v
}

// NewCreateBookPayload builds a books service create endpoint payload.
func NewCreateBookPayload(body *CreateRequestBody) *books.CreateBookPayload {
	v := &books.CreateBookPayload{
		Title:       *body.Title,
		Author:      *body.Author,
		Cover:       body.Cover,
		PublishedAt: body.PublishedAt,
	}

	return v
}

// NewUpdateBookPayload builds a books service update endpoint payload.
func NewUpdateBookPayload(body *UpdateRequestBody, id int) *books.UpdateBookPayload {
	v := &books.UpdateBookPayload{
		Title:       *body.Title,
		Author:      *body.Author,
		Cover:       body.Cover,
		PublishedAt: body.PublishedAt,
	}
	v.ID = id

	return v
}

// NewDeletePayload builds a books service delete endpoint payload.
func NewDeletePayload(id int) *books.DeletePayload {
	v := &books.DeletePayload{}
	v.ID = id

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.PublishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.published_at", *body.PublishedAt, goa.FormatDateTime))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.PublishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.published_at", *body.PublishedAt, goa.FormatDateTime))
	}
	return
}
