package main

//import (
//	"testing"
//
//	"github.com/aws/aws-lambda-go/events"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestHandler(t *testing.T) {
//	quesryStringMap := make(map[string]string)
//	quesryStringMap["username"] = "Naveen"
//	quesryStringMap["key"] = "123456"
//
//	inputHeaders := make(map[string]string)
//	inputHeaders["Authorization"] = "eyJraWQiOiIrT1loWmtJZTVZd2hQMEZFTUUxRzhnbzJLY1hXQzV0K01oNkNZOUNSTDI4PSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIyNDMzZWQwMi01MWJmLTQxYjMtODNmZC1kMTUxNDllMGM0MmEiLCJhdWQiOiI2NnRsY2hjdGs5djI4Y3Bhcm1ydXU1MjYxNSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZXZlbnRfaWQiOiI1YWRkOGE0OC02MjNmLTExZTgtYTJiYS05M2Q1OTBlZmM4MjYiLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTUyNzQ4ODQ4OCwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLXdlc3QtMi5hbWF6b25hd3MuY29tXC91cy13ZXN0LTJfMjhRdHRYTjZKIiwiY29nbml0bzp1c2VybmFtZSI6InRlc3R0ZXN0MTIzLWF0LWludHVpdC5jb20iLCJleHAiOjE1Mjc0OTYxNDMsImlhdCI6MTUyNzQ5MjU0NCwiZW1haWwiOiJ0ZXN0dGVzdDEyM0BpbnR1aXQuY29tIn0.VrgWLJ_n8u9cquRCvH9vU58DU4d9MWxWLRubXS3Q_iRsTyGiujZRu6r2NZ4iXhnSQpgwxl6kh_26KotXet3AL0pKP9UHKJA8lQVjQhnHDT9qDfrS6hQ2OptUFBkKzCKvVcudmoe8LcOfwszNFGgZ3iobf8fycqKXI5oaSL6Ef5GML3Qi_AYXZuEhY8Sg3JJWUdBdMcRqP5sWGna1OBMwpGrvd33m_PoVpVV0fmKUpxg0DavzPMnBlbdPmZhrQjDFLQHbRhy02jBjTxcD8XZT-NGNjqBdEDDZLBLaO46QehAYgSFEcd8N1uW7DisPJlJCIlBVCfI81CVlXg6vEM5VmQ"
//	tests := []struct {
//		request events.APIGatewayProxyRequest
//		expect  string
//		err     error
//	}{
//		{
//			// Test that the handler responds with the correct response
//			// when a valid name is provided in the HTTP body
//			request: events.APIGatewayProxyRequest{Path: "/keys", HTTPMethod: "POST", QueryStringParameters: quesryStringMap, Headers: inputHeaders, Body: "key=&value="},
//			expect:  "[{\"userName\":\"Naveen\",\"key\":\"123456\",\"value\":\"abcdefgh\"},{\"userName\":\"Naveen\",\"key\":\"23456\",\"value\":\"asdjhshf\"}]",
//			err:     nil,
//		},
//		{
//			// Test that the handler responds with the correct response
//			// when a valid name is provided in the HTTP body
//			request: events.APIGatewayProxyRequest{Path: "/keys", HTTPMethod: "GET", QueryStringParameters: quesryStringMap},
//			expect:  "[{\"userName\":\"Naveen\",\"key\":\"123456\",\"value\":\"abcdefgh\"},{\"userName\":\"Naveen\",\"key\":\"23456\",\"value\":\"asdjhshf\"}]",
//			err:     nil,
//		},
//		{
//			// Test that the handler responds with the correct response
//			// when a valid name is provided in the HTTP body
//			request: events.APIGatewayProxyRequest{Path: "/keys", HTTPMethod: "DELETE", QueryStringParameters: quesryStringMap},
//			expect:  "[{\"userName\":\"Naveen\",\"key\":\"123456\",\"value\":\"abcdefgh\"},{\"userName\":\"Naveen\",\"key\":\"23456\",\"value\":\"asdjhshf\"}]",
//			err:     nil,
//		},
//	}
//
//	for _, test := range tests {
//		response, err := Handler(test.request)
//		assert.IsType(t, test.err, err)
//		assert.Equal(t, test.expect, response.Body)
//	}
//
//}
