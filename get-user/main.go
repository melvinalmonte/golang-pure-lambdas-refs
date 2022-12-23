package main

import (
	"context"
	"encoding/json"
	"get-user-svc/services"
	"get-user-svc/utils"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func HandleRequest(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	zap.S().Info("Lambda called")

	uuid := event.PathParameters["uuid"]

	response, err := services.QueryToDynamo(uuid)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "An error has occurred querying user to DynamoDB, check logs for more details.", StatusCode: http.StatusBadRequest}, nil
	}

	body, _ := json.Marshal(response)

	event.Body = string(body)

	return events.APIGatewayProxyResponse{Body: event.Body, StatusCode: http.StatusOK}, nil
}

func main() {
	utils.InitLogger()
	lambda.Start(HandleRequest)
}
