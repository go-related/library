package database

import (
	"context"
	log "github.com/jt/books/gen/log"
	"github.com/jt/books/models"
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

type booksDatabase struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewBooksDB(connectionString string, logger *log.Logger) (BooksDB, error) {
	// note here we might want to pass our configured logger, so we don't log all the queries
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		logger.Error().Err(err).Msg("couldn't connect to the db")
		return nil, err
	}
	// Note this shouldn't be on a production code bc this will be executed everytime the application starts/restarts
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		logger.Warn().Err(err).Msg("couldn't migrate the db")
	}
	bookDb := booksDatabase{
		db:     db,
		logger: logger,
	}
	return &bookDb, err
}

func (b *booksDatabase) CreateBook(ctx context.Context, data models.Book) (models.Book, error) {
	result := b.db.Model(&models.Book{}).Create(&data)
	if result.Error != nil {
		b.logger.Error().Err(result.Error).Msg("couldn't create the book")
	}
	return data, result.Error
}

func (b *booksDatabase) UpdateBook(ctx context.Context, data models.Book) error {
	currentData, err := b.GetBookById(ctx, data.ID)
	if err != nil {
		// we dont log here since will be logged on getbyID
		return err
	}
	currentData.Title = data.Title
	currentData.Author = data.Author
	currentData.Cover = data.Cover
	currentData.PublishedAt = data.PublishedAt
	result := b.db.Save(currentData)
	if result.Error != nil {
		b.logger.Error().Err(result.Error).Uint("id", data.ID).Msg("couldn't update book details")
	}
	return result.Error
}

func (b *booksDatabase) DeleteBook(ctx context.Context, Id uint) error {
	currentData, err := b.GetBookById(ctx, Id)
	if err != nil {
		// we dont log here since will be logged on getbyID
		return err
	}
	result := b.db.Delete(&currentData)
	if result.Error != nil {
		b.logger.Error().Err(result.Error).Uint("id", Id).Msg("Error deleting book")
	}
	return nil
}

func (b *booksDatabase) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book
	result := b.db.Model(&models.Book{}).Find(&books)
	if result.Error != nil {
		b.logger.Error().Err(result.Error).Msg("couldn't load books")
	}
	return books, result.Error
}

func (b *booksDatabase) GetBookById(ctx context.Context, Id uint) (*models.Book, error) {
	var readBook models.Book
	result := b.db.Model(&models.Book{}).First(&readBook, Id)
	if result.Error != nil {
		b.logger.Error().Err(result.Error).Msg("couldn't create the book")
	}
	return &readBook, result.Error
}
