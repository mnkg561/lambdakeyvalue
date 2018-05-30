package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func putKeys(keyValue KeyValue) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	av, err := dynamodbattribute.MarshalMap(keyValue)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return err
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("KEY_VALUE_TABLE"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Successfully added key value to table")

	return nil
}

func createKeyValue(requestBody string, key string, value string) (string, string) {
	//key := uuid.Must(uuid.NewV4()).String()
	//value := uuid.Must(uuid.NewV4()).String()

	if requestBody != "" {
		formdata := strings.Split(requestBody, "&")
		tempKey := key
		tempValue := value
		for _, v := range formdata {
			param := strings.Split(v, "=")
			if param[1] != "" {
				if param[0] == "key" {
					key = param[1]
				}
				if param[0] == "value" {
					value = param[1]
				}
			}
		}

		if tempKey != key && tempValue == value {
			value = ""
		}
	}
	fmt.Printf("Key = %s \n", key)
	fmt.Printf("value = %s \n", value)
	return key, value
}
