package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getItems(userName string) (string, error) {
	var items []KeyValue
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	svc := dynamodb.New(sess)
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(userName),
			},
		},
		KeyConditionExpression: aws.String("userName = :v1"),
		TableName:              aws.String("KEY_VALUE_TABLE"),
	}

	result, err := svc.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		fmt.Println(err.Error())
		return "", err

	}

	keyValue := KeyValue{}

	for _, v := range result.Items {

		err = dynamodbattribute.UnmarshalMap(v, &keyValue)

		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		}

		items = append(items, keyValue)

	}

	response, err := json.Marshal(items)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
