package main

import (
	"context"
	"delete-user-svc/services"
	"delete-user-svc/utils"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func HandleRequest(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	zap.S().Info("Lambda called")

	uuid := event.PathParameters["uuid"]

	err := services.DeleteItemInDynamo(uuid)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "An error has occurred querying user to DynamoDB, check logs for more details.", StatusCode: http.StatusBadRequest}, nil
	}

	return events.APIGatewayProxyResponse{Body: "User ID successfully deleted", StatusCode: http.StatusOK}, nil
}

func main() {
	utils.InitLogger()
	lambda.Start(HandleRequest)
}
