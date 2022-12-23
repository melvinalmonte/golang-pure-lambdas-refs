package services

import (
	"context"
	"get-user-svc/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"
)

func QueryToDynamo(userId string) ([]models.UserModel, error) {
	zap.S().Info("-----Updating entry in DynamoDB request start-----")

	user := []models.UserModel{}

	zap.S().Info("Loading AWS config")
	awsConfig, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		zap.S().Errorw("An error has occurred initializing AWS config, error: %v", err)

	}

	zap.S().Info("Initializing DynamoDB service")
	svc := dynamodb.NewFromConfig(awsConfig)

	query := expression.Key("uuid").Equal(expression.Value(userId))
	expr, err := expression.NewBuilder().WithKeyCondition(query).Build()
	if err != nil {
		zap.S().Errorf("An error has occurred building query, error: %v", err)
	}

	res, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String("users"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		zap.S().Errorf("An error has occurred updating user entry, error: %v", err)

	}


	zap.S().Infof("Response item: %v", res.Items)
	err = attributevalue.UnmarshalListOfMaps(res.Items, &user)

	if err != nil {
		zap.S().Errorf("An error has occurred unmarshaling user query response, error: %v", err)

	}

	zap.S().Info("-----Updating entry in DynamoDB request update end-----")

	return user, err

}
