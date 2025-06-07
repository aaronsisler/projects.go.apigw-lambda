package dao

import (
	"context"
	"fmt"

	model "github.com/aaronsisler/projects.go.apigw-lambda/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserDao struct {
	db        *dynamodb.Client
	tableName string
}

func NewUserDao(db *dynamodb.Client, tableName string) *UserDao {
	return &UserDao{
		db:        db,
		tableName: tableName,
	}
}

func (dao *UserDao) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	result, err := dao.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(dao.tableName),
		Key: map[string]types.AttributeValue{
			"partitionKey": &types.AttributeValueMemberS{Value: "USER"},
			"sortKey":      &types.AttributeValueMemberS{Value: "USER#" + userId},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if result.Item == nil {
		return nil, nil
	}

	user := &model.User{
		UserId: userId,
	}

	if nameAttr, ok := result.Item["name"].(*types.AttributeValueMemberS); ok {
		user.Name = nameAttr.Value
	}

	if establishmentIdsAttr, ok := result.Item["establishmentIds"].(*types.AttributeValueMemberS); ok {
		user.EstablishmentIds = establishmentIdsAttr.Value
	}

	return user, nil
}
