package services

import (
	"bytes"
	"testing"
)

func TestUtilsEncoding(t *testing.T) {

	t.Run("CheckEncodingEmptySourceReturnsNil", func(t *testing.T) {

		//setup
		t.Log("Checking the EncodeBytesToBase64String returns nil on empty source")
		var toEncodeBytes []byte

		//execute
		result := EncodeBytesToBase64String(&toEncodeBytes)

		//assertion
		if result != nil {
			t.Fatal("failed to encode empty bytes")
		}

	})

	t.Run("CheckEncodingSourceReturnsCorrectResult", func(t *testing.T) {

		//setup
		t.Log("Checking the EncodeBytesToBase64String returns correct result")
		toEncodeBytes := []byte("Hello, world")
		expectedString := "SGVsbG8sIHdvcmxk"

		//execute
		result := EncodeBytesToBase64String(&toEncodeBytes)

		//assertion
		if result == nil {
			t.Fatal("failed to encode")
		}
		if *result != expectedString {
			t.Fatal("failed to encode empty bytes")
		}

	})
}

func TestUtilsDecoding(t *testing.T) {

	t.Run("CheckDecodingNilSourceReturnsNil", func(t *testing.T) {

		//setup
		t.Log("Checking the DecodeBase64StringToBytes returns nil on Nil source")
		var toDecodeString string

		//execute
		result, err := DecodeBase64StringToBytes(&toDecodeString)

		//assertion
		if err != nil {
			t.Fatal("failed to decode nil source")
		}
		if result != nil {
			t.Fatal("failed to decode nil source")
		}
	})

	t.Run("CheckDecodingEmptySourceReturnsNil", func(t *testing.T) {

		//setup
		t.Log("Checking the DecodeBase64StringToBytes returns nil on empty source")
		toDecodeString := ""

		//execute
		result, err := DecodeBase64StringToBytes(&toDecodeString)

		//assertion
		if err != nil {
			t.Fatal("failed to decode empty source")
		}
		if result != nil {
			t.Fatal("failed to decode empty source")
		}
	})

	t.Run("CheckDecodingNonEmptySourceReturnsCorrectResult", func(t *testing.T) {

		//setup
		t.Log("Checking the DecodeBase64StringToBytes returns correct result on non empty source")
		toDecodeString := "SGVsbG8sIHdvcmxk"
		expectedString := "Hello, world"

		//execute
		result, err := DecodeBase64StringToBytes(&toDecodeString)

		//assertion
		if err != nil {
			t.Fatal("failed to decode non empty source")
		}
		if !bytes.Equal([]byte(expectedString), *result) {
			t.Fatal("failed to decode non empty source")
		}
	})

}
