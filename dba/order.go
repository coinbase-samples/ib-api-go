package dba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/coinbase-samples/ib-api-go/model"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
)

const OrderIdIndex = "OrderIdGsi"

func (m *Repository) ListOrders(ctx context.Context, userId string) ([]model.Order, error) {
	var orders []model.Order

	l := ctxlogrus.Extract(ctx)
	l.Debugln("fetching order with ", userId)
	out, err := m.Svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(m.App.ActivityTableName),
		KeyConditionExpression: aws.String("userId = :a"),
		FilterExpression:       aws.String("orderStatus = :b"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":a": &types.AttributeValueMemberS{Value: userId},
			":b": &types.AttributeValueMemberS{Value: "open"},
		},
	})

	l.Debugln("query result for orders", out, err)

	if err != nil {
		return orders, err
	}

	l.Debugln(&out.Items)
	err = attributevalue.UnmarshalListOfMaps(out.Items, &orders)
	if err != nil {
		return orders, err
	}
	l.Debugln("unmarshalled orders", &orders)

	return orders, nil
}
