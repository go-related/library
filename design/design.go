package design

import (
	. "goa.design/goa/v3/dsl"
)
import _ "goa.design/plugins/v3/zerologger"

var _ = API("library", func() {
	Title("Library Service")
	Description("Service for managing library operations")
	Server("library", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("books", func() {

	Description("The book service performs CRUD operations on the books resource.")

	// error that will be returned when book is not found
	Error("not_found", func() {
		Description("Book not found")
	})
	// error that will be returned when custom validation has failed
	Error("bad_request", func() {
		Description("Invalid request")
	})

	// Crud operations
	Method("list", func() {
		Description("Retrieve all books")
		Result(ArrayOf(Book))

		HTTP(func() {
			GET("/books")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Retrieve a book by ID")
		Payload(func() {
			Field(1, "id", Int, "ID of the book")
			Required("id")
		})
		Result(Book)

		HTTP(func() {
			GET("/books/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("create", func() {
		Description("Create a new book")
		Payload(CreateBookPayload)
		Result(Book)

		HTTP(func() {
			POST("/books")
			Response(StatusCreated)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("update", func() {
		Description("Update an existing book")
		Payload(UpdateBookPayload)
		Result(Book)

		HTTP(func() {
			PUT("/books/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("delete", func() {
		Description("Delete a book by ID")
		Payload(func() {
			Field(1, "id", Int, "ID of the book")
			Required("id")
		})

		HTTP(func() {
			DELETE("/books/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
