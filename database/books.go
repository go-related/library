package database

import (
	"context"
	"github.com/jt/books/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BooksDB interface {
	CreateBook(ctx context.Context, data models.Book) (models.Book, error)
	UpdateBook(ctx context.Context, data models.Book) error
	DeleteBook(ctx context.Context, Id uint) error
	GetAllBooks(ctx context.Context) ([]*models.Book, error)
	GetBookById(ctx context.Context, Id uint) (*models.Book, error)
}

type booksDb struct {
	db *gorm.DB
}

func (b booksDb) CreateBook(ctx context.Context, data models.Book) (models.Book, error) {
	result := b.db.Model(&models.Book{}).Create(&data)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't create the book")
	}
	return data, result.Error
}

func (b booksDb) UpdateBook(ctx context.Context, data models.Book) error {
	currentData, err := b.GetBookById(ctx, data.ID)
	if err != nil {
		return err
	}
	currentData.Title = data.Title
	currentData.Author = data.Author
	currentData.Cover = data.Cover
	currentData.PublishedAt = data.PublishedAt
	result := b.db.Save(currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", data.ID).Error("couldn't update book details")
	}
	return result.Error
}

func (b booksDb) DeleteBook(ctx context.Context, Id uint) error {
	currentData, err := b.GetBookById(ctx, Id)
	if err != nil {
		return err
	}
	result := b.db.Delete(&currentData)
	if result.Error != nil {
		logrus.WithError(result.Error).WithField("id", Id).Error("Error deleting book")
	}
	return nil
}

func (b booksDb) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book
	result := b.db.Model(&models.Book{}).Find(&books)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't load books")
	}
	return books, result.Error
}

func (b booksDb) GetBookById(ctx context.Context, Id uint) (*models.Book, error) {
	var readBook models.Book
	result := b.db.Model(&models.Book{}).First(&readBook, Id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("couldn't create the book")
	}
	return &readBook, result.Error
}

func NewBooksDB(connectionString string) (BooksDB, error) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Error("couldn't connect to the db")
		return nil, err
	}
	// Note this shouldn't be on a production code bc this will be executed everytime the application starts/restarts
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		logrus.WithError(err).Warn("couldn't migrate the db")
	}
	bookDb := booksDb{
		db: db,
	}
	return &bookDb, err
}
