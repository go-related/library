package services

import (
	b64 "encoding/base64"
)

func EncodeBytesToBase64String(sc *[]byte) *string {
	if len(*sc) == 0 {
		return nil
	}
	result := b64.StdEncoding.EncodeToString(*sc)
	return &result
}

func DecodeBase64StringToBytes(sc *string) (*[]byte, error) {
	if len(*sc) == 0 {
		return nil, nil
	}
	result, err := b64.StdEncoding.DecodeString(*sc)
	return &result, err
}
