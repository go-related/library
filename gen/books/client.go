// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books client
//
// Command:
// $ goa gen github.com/jt/books/design

package books

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "books" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	ShowEndpoint   goa.Endpoint
	CreateEndpoint goa.Endpoint
	UpdateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
}

// NewClient initializes a "books" service client given the endpoints.
func NewClient(list, show, create, update, delete_ goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		ShowEndpoint:   show,
		CreateEndpoint: create,
		UpdateEndpoint: update,
		DeleteEndpoint: delete_,
	}
}

// List calls the "list" endpoint of the "books" service.
// List may return the following errors:
//   - "not_found" (type *goa.ServiceError): Book not found
//   - "bad_request" (type *goa.ServiceError): Invalid request
//   - error: internal error
func (c *Client) List(ctx context.Context) (res []*Book, err error) {
	var ires any
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Book), nil
}

// Show calls the "show" endpoint of the "books" service.
// Show may return the following errors:
//   - "not_found" (type *goa.ServiceError): Book not found
//   - "bad_request" (type *goa.ServiceError): Invalid request
//   - error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *Book, err error) {
	var ires any
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Book), nil
}

// Create calls the "create" endpoint of the "books" service.
// Create may return the following errors:
//   - "not_found" (type *goa.ServiceError): Book not found
//   - "bad_request" (type *goa.ServiceError): Invalid request
//   - error: internal error
func (c *Client) Create(ctx context.Context, p *CreateBookPayload) (res *Book, err error) {
	var ires any
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Book), nil
}

// Update calls the "update" endpoint of the "books" service.
// Update may return the following errors:
//   - "not_found" (type *goa.ServiceError): Book not found
//   - "bad_request" (type *goa.ServiceError): Invalid request
//   - error: internal error
func (c *Client) Update(ctx context.Context, p *UpdateBookPayload) (res *Book, err error) {
	var ires any
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Book), nil
}

// Delete calls the "delete" endpoint of the "books" service.
// Delete may return the following errors:
//   - "not_found" (type *goa.ServiceError): Book not found
//   - "bad_request" (type *goa.ServiceError): Invalid request
//   - error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}
