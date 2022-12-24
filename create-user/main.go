package main

import (
	"context"
	"create-user-svc/models"
	"create-user-svc/services"
	"create-user-svc/utils"
	"encoding/json"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func HandleRequest(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	zap.S().Info("Lambda called")
	zap.S().Infof("Event body: %v", event.Body)

	var userPayload models.UserModel

	err := json.Unmarshal([]byte(event.Body), &userPayload)

	if err != nil {
		zap.S().Errorw("An error has occurred marshaling the request body, error: %v", err)
		return events.APIGatewayProxyResponse{Body: "An error has occurred, check logs for more details.", StatusCode: http.StatusBadRequest}, nil
	}

	response, err := services.WriteToDynamo(userPayload)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "An error has occurred writing to DynamoDB, check logs for more details.", StatusCode: http.StatusBadRequest}, nil
	}

	body, _ := json.Marshal(response)

	event.Body = string(body)

	return events.APIGatewayProxyResponse{Body: event.Body, StatusCode: http.StatusOK}, nil
}

func main() {
	utils.InitLogger()
	lambda.Start(HandleRequest)
}
