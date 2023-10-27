package services_test

import (
	"context"
	"github.com/jt/books/gen/books"
	log "github.com/jt/books/gen/log"
	"github.com/jt/books/models"
	"github.com/jt/books/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBookHappyPath(t *testing.T) {
	logger := log.New("book_tests", false)
	mockDB := NewMockedBooksDb()
	service := services.NewBooks(logger, mockDB)
	t.Run("BooksListReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Testing that method list of the service bookssrvc returns correct result")
		expectedResult := []*models.Book{
			NewBook("book1", "author1", 1),
			NewBook("book2", "author2", 2),
		}
		mockDB.On("GetAllBooks", context.Background()).Return(expectedResult, nil)

		//execute
		result, err := service.List(context.Background())

		//assertion
		assert.Nil(t, err) // No error should be returned
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(expectedResult))
	})

	t.Run("BooksShowReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'show' of the service bookssrvc returns correct result")
		expectedBook := NewBook("book2", "author2", 2)
		mockDB.On("GetBookById", context.Background(), uint(2)).Return(expectedBook, nil)

		//execute
		result, err := service.Show(context.Background(), &books.ShowPayload{ID: 2})

		//assertion
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedBook.Title, *result.Title)
		assert.Equal(t, expectedBook.Author, *result.Author)
		assert.Equal(t, int(expectedBook.ID), *result.ID)
	})

	t.Run("BooksDeleteReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'delete' of the service bookssrvc returns correct result")
		mockDB.On("DeleteBook", context.Background(), uint(2)).Return(nil)

		//execute
		err := service.Delete(context.Background(), &books.DeletePayload{ID: 2})

		//assertion
		assert.Nil(t, err)
	})

	t.Run("BooksCreateReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'create' of the service bookssrvc returns correct result")

		payload := books.CreateBookPayload{
			Title:  "title",
			Author: "test",
		}
		expectedBook := NewBook("title", "test", 3)
		mockDB.On("CreateBook", context.Background(), mock.AnythingOfType("models.Book")).Return(*expectedBook, nil)

		//execute
		result, err := service.Create(context.Background(), &payload)

		//assertion
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedBook.Title, *result.Title)
		assert.Equal(t, expectedBook.Author, *result.Author)
		assert.Equal(t, int(expectedBook.ID), *result.ID)

	})

	t.Run("BooksUpdateReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'update' of the service bookssrvc returns correct result")
		payload := books.UpdateBookPayload{
			Title:  "title",
			Author: "test",
			ID:     6,
		}
		expectedBook := NewBook("title", "test", 6)
		mockDB.On("UpdateBook", context.Background(), mock.AnythingOfType("models.Book")).Return(nil)

		//execute
		result, err := service.Update(context.Background(), &payload)

		//assertion
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedBook.Title, *result.Title)
		assert.Equal(t, expectedBook.Author, *result.Author)
		assert.Equal(t, int(expectedBook.ID), *result.ID)
	})
}

//func TestBookUnHappyPath(t *testing.T) {
//
//	t.Run("CheckConvertingEmptyBookModelReturnsEmptyResult", func(t *testing.T) {
//		//setup
//		t.Log("Checking that method convertBooksToGoaModel returns nil on empty source")
//		var toConvertModel models.Book
//
//		//execute
//		result := convertBooksToGoaModel(&toConvertModel)
//
//		//assertion
//		if result != nil {
//			t.Fatal("failed to return proper response on empty model")
//		}
//
//	})
//}

func NewBook(title, author string, id uint) *models.Book {
	result := models.Book{
		Title:  title,
		Author: author,
	}
	result.ID = id
	return &result
}

type MockBooksDB struct {
	mock.Mock
}

func NewMockedBooksDb() *MockBooksDB {
	return &MockBooksDB{}
}

func (m *MockBooksDB) CreateBook(ctx context.Context, data models.Book) (models.Book, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(models.Book), args.Error(1)
}

func (m *MockBooksDB) UpdateBook(ctx context.Context, data models.Book) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *MockBooksDB) DeleteBook(ctx context.Context, Id uint) error {
	args := m.Called(ctx, Id)
	return args.Error(0)
}

func (m *MockBooksDB) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Book), args.Error(1)
}

func (m *MockBooksDB) GetBookById(ctx context.Context, Id uint) (*models.Book, error) {
	args := m.Called(ctx, Id)
	return args.Get(0).(*models.Book), args.Error(1)
}
