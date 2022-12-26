package services

import (
	"context"
	"update-user-svc/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"go.uber.org/zap"
)

func UpdateItemInDynamo(user models.UserModel, userId string) error {
	zap.S().Info("-----Updating entry in DynamoDB request start-----")

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

	zap.S().Info("Marshaling attribute values (av) for user object")
	av, err := attributevalue.MarshalMap(user)

	if err != nil {
		zap.S().Errorf("An error has occurred marshaling user object, error: %v", err)
	}

	update := expression.UpdateBuilder{}

	for k, v := range av {
		update = update.Set(expression.Name(k), expression.Value(v))
	}

	zap.S().Info("Building DynamoDB update expression")
	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		zap.S().Errorw("An error has occurred build expression for update, error: %v", err)
	}

	_, err = svc.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String("users"),
		Key:                       map[string]types.AttributeValue{"uuid": id},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})

	if err != nil {
		zap.S().Errorf("An error has occurred updating user entry, error: %v", err)

	}

	zap.S().Info("-----Updating entry in DynamoDB request update end-----")

	return err

}
