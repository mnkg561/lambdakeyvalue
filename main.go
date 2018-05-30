package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/satori/go.uuid"
	"strings"
)

type customError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type customResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

//Handler process APIGatewayProxyRequest event and based on HTTP Method in the input,
//request will be processed and will respond in APIGatewYProxyResponse event
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	fmt.Println("Received method: ", request.HTTPMethod)
	fmt.Println("Received header: ", request.Headers)
	fmt.Println("Received path:", request.Path)
	fmt.Println("Receieved cognito user ", request.RequestContext.Authorizer)
	fmt.Println("extra comment")

	headersMap := make(map[string]string)
	headersMap["Access-Control-Allow-Origin"] = "*"

	if request.Path == "/v1/healthCheck" {

		return events.APIGatewayProxyResponse{Body: "GREEN", StatusCode: 200, Headers: headersMap}, nil
	}

	var userInfo UserInfo

	incomingHeadersMap := request.Headers

	authToken := incomingHeadersMap["Authorization"]
	tokens := strings.Split(authToken, ".")
	userPart := tokens[1]

	sDec, err := b64.RawStdEncoding.DecodeString(userPart)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(sDec, &userInfo)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("user id in the request %s", userInfo.Email)
	fmt.Println()

	userName := userInfo.Email

	error1 := &customError{Code: "400", Message: "BadRequest", Detail: "Invalid path"}
	response, err := json.Marshal(error1)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if request.HTTPMethod == "GET" {

		response, err := getItems(userName)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		return events.APIGatewayProxyResponse{Body: response, StatusCode: 200, Headers: headersMap}, nil

	} else if request.HTTPMethod == "POST" {

		key := uuid.Must(uuid.NewV4()).String()
		value := uuid.Must(uuid.NewV4()).String()

		key, value = createKeyValue(request.Body, key, value)
		putKeys(KeyValue{UserName: userName, Key: key, Value: value})

		response1 := &customResponse{Code: "200", Message: "Success", Detail: "Key and Value Successfully created"}
		response, err = json.Marshal(response1)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	} else if request.HTTPMethod == "DELETE" {
		key := "123455"
		deleteKeys(userName, key)
		response1 := &customResponse{Code: "200", Message: "Success", Detail: "Key got deleted Successfully"}
		response, err = json.Marshal(response1)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	}

	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200, Headers: headersMap}, nil
}

func main() {
	lambda.Start(Handler)
}
