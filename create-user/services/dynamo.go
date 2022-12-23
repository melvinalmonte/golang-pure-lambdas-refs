package services

import (
	"context"
	"create-user-svc/models"

	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"

	"github.com/google/uuid"
)

func WriteToDynamo(user models.UserModel) (models.UserModel, error) {
	zap.S().Info("-----Writing user to DynamoDB request start-----")

	zap.S().Info("Loading AWS config")
	awsConfig, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		zap.S().Errorw("An error has occurred initializing AWS config, error: %v", err)

	}

	zap.S().Info("Initializing DynamoDB service")
	svc := dynamodb.NewFromConfig(awsConfig)

	zap.S().Info("Generating UUID for new user record")
	id, err := uuid.NewUUID()

	if err != nil {
		zap.S().Errorf("An error has occurred generating UUID, error: %v", err)
	}

	currentTime := time.Now()

	user.Uuid = id.String()
	user.CreatedEntry = currentTime.UTC().String()

	record, err := attributevalue.MarshalMap(user)
	if err != nil {
		zap.S().Errorf("An error has occurred marshaling user object, error: %v", err)
	}

	zap.S().Infof("Adding %v to users table", user)
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("users"),
		Item:      record,
	})

	if err != nil {
		zap.S().Errorf("An error has occurred writing user object to DynamoDB, error: %v", err)
	}

	zap.S().Info("-----Writing user to DynamoDB request end-----")
	return user, err

}
