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
		toConvertDate, _ := time.Parse(time.RFC3339, "2023-10-27T09:10:14Z")
		toConvertId := 1
		toConvertModel := models.Book{
			Title:       "testTitle",
			Author:      "",
			Cover:       nil,
			PublishedAt: &toConvertDate,
		}
		toConvertModel.Model.ID = uint(toConvertId)

		expectedDate := "2023-10-27T09:10:14Z"
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
		compareResult, compareError := CompareStructFields(expected, result)
		if !compareResult {
			t.Fatalf("failed to return proper response (%s)", compareError)
		}

	})

	t.Run("CheckConvertingNonEmptyCreateBookPayloadReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Checking that method convertCreateBookPayloadToBookModel returns correct result")
		toConvertDate, _ := time.Parse(time.RFC3339, "2023-10-27T09:10:14Z")
		expectedDate := "2023-10-27T09:10:14Z"
		toConvertCover := ""
		toConvertModel := books.CreateBookPayload{
			Author:      "test",
			Title:       "testTitle",
			PublishedAt: &expectedDate,
			Cover:       &toConvertCover,
		}
		expectedCover := []byte(toConvertCover)
		expected := models.Book{
			PublishedAt: &toConvertDate,
			Author:      "test",
			Title:       "testTitle",
			Cover:       &expectedCover,
		}

		//execute
		result := convertCreateBookPayloadToBookModel(&toConvertModel, &expectedCover)

		//assertion
		compareResult, compareError := CompareStructFields(expected, result)
		if !compareResult {
			t.Fatalf("failed to return proper response (%s)", compareError)
		}

	})

	t.Run("CheckConvertingEmptyUpdateBookPayloadReturnsEmptyResult", func(t *testing.T) {
		//setup
		t.Log("Checking that method convertUpdatePayloadToBookModel returns nil on empty source")
		var toConvertModel books.UpdateBookPayload
		var expectedDate time.Time
		expected := models.Book{
			PublishedAt: &expectedDate,
		}

		//execute
		result := convertUpdatePayloadToBookModel(&toConvertModel, nil)

		//assertion
		compareResult, compareError := CompareStructFields(expected, result)
		if !compareResult {
			t.Fatalf("failed to return proper response (%s)", compareError)
		}

	})

	t.Run("CheckConvertingNonEmptyUpdateBookPayloadReturnsCorrectResult", func(t *testing.T) {
		//setup
		t.Log("Checking that method convertUpdatePayloadToBookModel returns correct result")
		toConvertDate, _ := time.Parse(time.RFC3339, "2023-10-27T09:10:14Z")
		expectedDate := "2023-10-27T09:10:14Z"
		toConvertCover := "testing"
		toConvertModel := books.UpdateBookPayload{
			Author:      "test",
			Title:       "testTitle",
			PublishedAt: &expectedDate,
			Cover:       &toConvertCover,
			ID:          10,
		}
		expectedCover := []byte(toConvertCover)
		expected := models.Book{
			PublishedAt: &toConvertDate,
			Author:      "test",
			Title:       "testTitle",
			Cover:       &expectedCover,
		}
		expected.ID = 10

		//execute
		result := convertUpdatePayloadToBookModel(&toConvertModel, &expectedCover)

		//assertion
		compareResult, compareError := CompareStructFields(expected, result)
		if !compareResult {
			t.Fatalf("failed to return proper response (%s)", compareError)
		}

	})
}

func CompareStructFields(expected, result interface{}) (bool, string) {
	expectedValues := reflect.ValueOf(expected)
	resultValues := reflect.ValueOf(result)
	t := reflect.TypeOf(expected)
	for i := 0; i < expectedValues.NumField(); i++ {
		f1 := expectedValues.Field(i).Interface()
		f2 := resultValues.Field(i).Interface()
		resultValues.Field(i).Type()
		if !reflect.DeepEqual(f1, f2) {
			return false, fmt.Sprintf("field %v is not equal", t.Field(i).Name)
		}
	}
	return true, ""
}
