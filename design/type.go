package design

import (
	. "goa.design/goa/v3/dsl"
)

// Book represents a book entity.
var Book = Type("Book", func() {
	Description("book entity.")
	Attribute("id", Int, "ID of the book")
	Attribute("title", String, "Title of the book")
	Attribute("author", String, "Author of the book")
	Attribute("cover", String, func() {
		Description("Base64 of the Book cover image")
		Example("QmFzZTY0IG9mIHRoZSBCb29rIGNvdmVyIGltYWdl")
	})
	Attribute("published_at", String, func() {
		Description("Published date of the book")
		Format(FormatDateTime)
		Example("2006-01-02T15:04:05Z")
	})
})

// CreateBookPayload payload for creating a new book
var CreateBookPayload = Type("CreateBookPayload", func() {
	Field(1, "title", String, "Title of the book")
	Field(2, "author", String, "Author of the book")
	Field(3, "cover", String, func() {
		Description("Base64 of the Book cover image")
		Example("QmFzZTY0IG9mIHRoZSBCb29rIGNvdmVyIGltYWdl")
	})
	Field(4, "published_at", String, func() {
		Description("Published date of the book")
		Format(FormatDateTime)
		Example("2006-01-02T15:04:05Z")
	})
	Required("title", "author")
})

// UpdateBookPayload payload for creating a new book
var UpdateBookPayload = Type("UpdateBookPayload", func() {
	Field(1, "id", Int, "ID of the book")
	Field(2, "title", String, "Title of the book")
	Field(3, "author", String, "Author of the book")
	Field(4, "cover", String, func() {
		Description("Base64 of the Book cover image")
		Example("QmFzZTY0IG9mIHRoZSBCb29rIGNvdmVyIGltYWdl")
	})
	Field(5, "published_at", String, func() {
		Description("Published date of the book")
		Format(FormatDateTime)
		Example("2006-01-02T15:04:05Z")
	})
	Required("id", "title", "author")
})
