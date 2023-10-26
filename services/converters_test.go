package services

import (
	"fmt"
	"github.com/jt/books/gen/books"
	"github.com/jt/books/models"
	"reflect"
	"testing"
	"time"
)

func TestConverters(t *testing.T) {

	t.Run("CheckConvertingEmptyBookModelReturnsEmptyResult", func(t *testing.T) {
		//setup
		t.Log("Checking that method convertBooksToGoaModel returns nil on empty source")
		var toConvertModel models.Book

		//execute
		result := convertBooksToGoaModel(&toConvertModel)

		//assertion
		if result != nil {
			t.Fatal("failed to return proper response on empty model")
		}

	})

	t.Run("CheckConvertingBookModelReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Checking that method convertBooksToGoaModel returns correct result")
		toConvertDate := time.Now()
		toConvertId := 1
		toConvertModel := models.Book{
			Title:       "testTitle",
			Author:      "",
			Cover:       nil,
			PublishedAt: &toConvertDate,
		}
		toConvertModel.Model.ID = uint(toConvertId)

		expectedDate := toConvertDate.Format(time.RFC3339)
		expectedData := books.Book{
			ID:          &toConvertId,
			Title:       &toConvertModel.Title,
			Author:      &toConvertModel.Author,
			PublishedAt: &expectedDate,
			Cover:       nil,
		}
		//execute
		result := convertBooksToGoaModel(&toConvertModel)

		//assertion
		if result == nil {
			t.Fatal("failed to return proper response on non empty model")
		}

		if *result.ID != *expectedData.ID {
			t.Fatal("invalid id when converting model")
		}
		if result.Cover != expectedData.Cover {
			t.Fatal("invalid cover when converting model")
		}
		if result.Title != expectedData.Title {
			t.Fatal("invalid title when converting model")
		}
		if result.Author != expectedData.Author {
			t.Fatal("invalid author when converting model")
		}
		if *result.PublishedAt != *expectedData.PublishedAt {
			t.Fatal("invalid published_at when converting model")
		}
	})

	t.Run("CheckConvertingEmptyCreateBookPayloadReturnsEmptyResult", func(t *testing.T) {
		//setup
		t.Log("Checking that method convertCreateBookPayloadToBookModel returns nil on empty source")
		var toConvertModel books.CreateBookPayload
		var expectedDate time.Time
		expected := models.Book{
			PublishedAt: &expectedDate,
		}

		//execute
		result := convertCreateBookPayloadToBookModel(&toConvertModel, nil)

		//assertion
		compareResult, compareError := compareStructFields(expected, result)
		if !compareResult {
			t.Fatalf("failed to return proper response (%s)", compareError)
		}

	})

	//TODO non empty result
	t.Run("CheckConvertingEmptyCreateBookPayloadReturnsEmptyResult", func(t *testing.T) {
		//setup
		t.Log("Checking that method convertCreateBookPayloadToBookModel returns nil on empty source")
		var toConvertModel books.CreateBookPayload
		var expectedDate time.Time
		expected := models.Book{
			PublishedAt: &expectedDate,
		}

		//execute
		result := convertCreateBookPayloadToBookModel(&toConvertModel, nil)

		//assertion
		compareResult, compareError := compareStructFields(expected, result)
		if !compareResult {
			t.Fatalf("failed to return proper response (%s)", compareError)
		}

	})
}

func compareStructFields(expected, result interface{}) (bool, string) {
	expectedValues := reflect.ValueOf(expected)
	resultValues := reflect.ValueOf(result)
	t := reflect.TypeOf(expected)
	for i := 0; i < expectedValues.NumField(); i++ {
		f1 := expectedValues.Field(i).Interface()
		f2 := resultValues.Field(i).Interface()
		if !reflect.DeepEqual(f1, f2) {
			return false, fmt.Sprintf("field %v is not equal", t.Field(i).Name)
		}
	}
	return true, ""
}
