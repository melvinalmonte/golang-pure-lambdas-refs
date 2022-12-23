package main

import (
	"context"
	"encoding/json"
	"update-user-svc/models"
	"update-user-svc/services"
	"update-user-svc/utils"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func HandleRequest(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	zap.S().Info("Lambda called")

	var userPayload models.UserModel
	uuid := event.PathParameters["uuid"]

	json.Unmarshal([]byte(event.Body), &userPayload)

	err := services.UpdateItemInDynamo(userPayload, uuid)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "An error has occurred writing to DynamoDB, check logs for more details.", StatusCode: http.StatusBadRequest}, nil
	}

	return events.APIGatewayProxyResponse{Body: "Entry has been successfully updated.", StatusCode: http.StatusOK}, nil
}

func main() {
	utils.InitLogger()
	lambda.Start(HandleRequest)
}
