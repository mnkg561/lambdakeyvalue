package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyValueGen(t *testing.T) {
	tests := []struct {
		requestBody    string
		generatedKey   string
		generatedValue string
		expectedKey    string
		expectedValue  string
	}{
		{
			// Empty key value pairs.. this will be part of the request when user clicks on "Generate Keyvalue" button
			requestBody:    "key=testkey123&value=testvalue123456",
			generatedKey:   "36b7b4db-90c2-47e9-a07c-099738a80083",
			generatedValue: "adabd2aa-828b-4a63-828f-bf3941587326",
			expectedKey:    "testkey123",
			expectedValue:  "testvalue123456",
		},
		{
			// Empty key value pairs.. this will be part of the request when user clicks on "Generate Keyvalue" button
			requestBody:    "key=&value=",
			generatedKey:   "36b7b4db-90c2-47e9-a07c-099738a80083",
			generatedValue: "adabd2aa-828b-4a63-828f-bf3941587326",
			expectedKey:    "36b7b4db-90c2-47e9-a07c-099738a80083",
			expectedValue:  "adabd2aa-828b-4a63-828f-bf3941587326",
		},
		{
			// Empty key value pairs.. this will be part of the request when user clicks on "Generate Keyvalue" button
			requestBody:    "key=123123&value=",
			generatedKey:   "36b7b4db-90c2-47e9-a07c-099738a80083",
			generatedValue: "adabd2aa-828b-4a63-828f-bf3941587326",
			expectedKey:    "123123",
			expectedValue:  "",
		},
	}

	for _, test := range tests {
		key, value := createKeyValue(test.requestBody, test.generatedKey, test.generatedValue)
		//assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedKey, key)
		assert.Equal(t, test.expectedValue, value)
	}

}
