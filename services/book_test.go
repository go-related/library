package services_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/jt/books/gen/books"
	log "github.com/jt/books/gen/log"
	"github.com/jt/books/models"
	"github.com/jt/books/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
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

func TestBookUnHappyPath(t *testing.T) {
	logger := log.New("book_tests", false)
	mockDB := NewMockedBooksDb()
	service := services.NewBooks(logger, mockDB)
	t.Parallel()
	t.Run("BooksListReturnsCorrectError", func(t *testing.T) {
		//setup
		t.Log("Testing that method list of the service bookssrvc returns the proper error")
		expectedError := fmt.Errorf("some error happened")
		mockDB.On("GetAllBooks", context.Background()).Return([]*models.Book{}, expectedError)

		//execute
		_, err := service.List(context.Background())

		//assertion
		assert.NotNil(t, err)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("BooksShowReturnsNotFoundError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'show' of the service bookssrvc returns not found error")
		expectedError := gorm.ErrRecordNotFound
		mockDB.On("GetBookById", context.Background(), uint(2)).Return(&models.Book{}, expectedError)

		//execute
		result, err := service.Show(context.Background(), &books.ShowPayload{ID: 2})

		//assertion
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Error(t, expectedError, err)
	})

	t.Run("BooksShowReturnsError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'show' of the service bookssrvc returns other errors")
		expectedError := fmt.Errorf("some error happened")
		mockDB.On("GetBookById", context.Background(), uint(2)).Return(&models.Book{}, expectedError)

		//execute
		result, err := service.Show(context.Background(), &books.ShowPayload{ID: 2})

		//assertion
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Error(t, expectedError, err)
	})

	t.Run("BooksDeleteReturnsError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'delete' of the service bookssrvc returns other errors")
		expectedError := fmt.Errorf("some error happened")
		mockDB.On("DeleteBook", context.Background(), uint(2)).Return(expectedError)

		//execute
		err := service.Delete(context.Background(), &books.DeletePayload{ID: 2})

		//assertion
		assert.NotNil(t, err)
		assert.Error(t, expectedError, err)
	})

	t.Run("BooksDeleteReturnsNotFoundError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'delete' of the service bookssrvc returns not found errors")
		expectedError := gorm.ErrRecordNotFound
		mockDB.On("DeleteBook", context.Background(), uint(2)).Return(expectedError)

		//execute
		err := service.Delete(context.Background(), &books.DeletePayload{ID: 2})

		//assertion
		assert.NotNil(t, err)
		assert.Error(t, expectedError, err)
	})

	t.Run("BooksCreateReturnsBadRequestError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'create' of the service bookssrvc returns bad request error on invalid base64 cover")
		expectedError := books.MakeBadRequest(fmt.Errorf("invalid string error"))
		invalidCover := "invalid_base64_string"
		payload := books.CreateBookPayload{
			Title:  "title",
			Author: "test",
			Cover:  &invalidCover,
		}
		expectedBook := NewBook("title", "test", 3)
		mockDB.On("CreateBook", context.Background(), mock.AnythingOfType("models.Book")).Return(*expectedBook, expectedError)

		//execute
		_, err := service.Create(context.Background(), &payload)

		//assertion
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expectedError))
	})

	t.Run("BooksCreateReturnsDbError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'create' of the service bookssrvc returns  error on case a db error happened")
		expectedError := fmt.Errorf("invalid string error")
		payload := books.CreateBookPayload{
			Title:  "title",
			Author: "test",
		}
		mockDB.On("CreateBook", context.Background(), mock.AnythingOfType("models.Book")).Return(models.Book{}, expectedError)

		//execute
		_, err := service.Create(context.Background(), &payload)

		//assertion
		assert.NotNil(t, err)
		assert.True(t, err.Error() == expectedError.Error())
	})

	t.Run("BooksUpdateReturnsBadRequestError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'update' of the service bookssrvc returns bad request error on invalid base64 cover")
		expectedError := books.MakeBadRequest(fmt.Errorf("invalid string error"))
		invalidCover := "invalid_base64_string"
		payload := books.UpdateBookPayload{
			Title:  "title",
			Author: "test",
			Cover:  &invalidCover,
		}
		mockDB.On("UpdateBook", context.Background(), mock.AnythingOfType("models.Book")).Return(expectedError)

		//execute
		_, err := service.Update(context.Background(), &payload)

		//assertion
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expectedError))
	})

	t.Run("BooksUpdateReturnsNotFoundError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'update' of the service bookssrvc returns not found")
		expectedError := gorm.ErrRecordNotFound
		payload := books.UpdateBookPayload{
			Title:  "title",
			Author: "test",
		}
		mockDB.On("UpdateBook", context.Background(), mock.AnythingOfType("models.Book")).Return(expectedError)

		//execute
		_, err := service.Update(context.Background(), &payload)

		//assertion
		assert.NotNil(t, err)
		assert.Error(t, expectedError, err)
	})

	t.Run("BooksUpdateReturnsDbError", func(t *testing.T) {
		//setup
		t.Log("Testing that method 'update' of the service bookssrvc returns all kinds of error")
		expectedError := fmt.Errorf("invalid string error")
		payload := books.UpdateBookPayload{
			Title:  "title",
			Author: "test",
			ID:     6,
		}
		mockDB.On("UpdateBook", context.Background(), mock.AnythingOfType("models.Book")).Return(expectedError)

		//execute
		_, err := service.Update(context.Background(), &payload)

		//assertion
		assert.NotNil(t, err)
		assert.Error(t, expectedError, err)
	})
}

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
