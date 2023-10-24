package design

import (
	. "goa.design/goa/v3/dsl"
)

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

// Book represents a book entity.
var Book = Type("Book", func() {
	Description("book entity.")
	Attribute("id", Int, "ID of the book")
	Attribute("title", String, "Title of the book")
	Attribute("author", String, "Author of the book")
	Attribute("cover", String, "Book cover image URL")
	Attribute("published_at", String, "Published date of the book")
})

// CreateBookPayload payload for creating a new book
var CreateBookPayload = Type("CreateBookPayload", func() {
	Field(1, "title", String, "Title of the book")
	Field(2, "author", String, "Author of the book")
	Field(3, "cover", String, "Book cover image URL")
	Field(4, "published_at", String, "Published date of the book")
	Required("title", "author")
})

// UpdateBookPayload payload for creating a new book
var UpdateBookPayload = Type("UpdateBookPayload", func() {
	Field(1, "id", Int, "ID of the book")
	Field(2, "title", String, "Title of the book")
	Field(3, "author", String, "Author of the book")
	Field(4, "cover", String, "Book cover image URL")
	Field(5, "published_at", String, "Published date of the book")
	Required("id", "title", "author")
})

var _ = Service("books", func() {
	Description("The book service performs CRUD operations on the books resource.")

	// error that will be returned when book is not found
	Error("not_found", func() {
		Description("Book not found")
	})

	// Crud operations
	Method("items", func() {
		Description("Retrieve all books")
		Result(ArrayOf(Book))

		HTTP(func() {
			GET("/books")
			Response(StatusOK)
		})
	})

	Method("item", func() {
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
