package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"go.uber.org/zap"
)

func DeleteItemInDynamo(userId string) error {
	zap.S().Info("-----Deleting entry in DynamoDB request start-----")

	zap.S().Info("Loading AWS config")
	awsConfig, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		zap.S().Errorw("An error has occurred initializing AWS config, error: %v", err)

	}

	zap.S().Info("Initializing DynamoDB service")
	svc := dynamodb.NewFromConfig(awsConfig)

	id, err := attributevalue.Marshal(userId)

	if err != nil {
		zap.S().Errorf("An error has occurred marshaling user ID, error: %v", err)
	}

	_, err = svc.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("users"),
		Key:       map[string]types.AttributeValue{"uuid": id},
	})

	if err != nil {
		zap.S().Errorf("An error has occurred deleting user entry, error: %v", err)

	}

	zap.S().Info("-----Deleting entry in DynamoDB request end-----")

	return err

}
