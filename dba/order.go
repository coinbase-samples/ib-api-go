package dba

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/coinbase-samples/ib-api-go/model"
)

const OrderIdIndex = "OrderIdGsi"

func (m *DynamoRepository) ListOrders(ctx context.Context, userId string) ([]model.Order, error) {
	var orders []model.Order

	log.DebugCtx(ctx, "fetching order with ", userId)
	out, err := m.Svc.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(m.App.ActivityTableName),
		KeyConditionExpression: aws.String("userId = :a"),
		FilterExpression:       aws.String("orderStatus = :b OR orderStatus = :c OR orderStatus = :d"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":a": &types.AttributeValueMemberS{Value: userId},
			":b": &types.AttributeValueMemberS{Value: "open"},
			":c": &types.AttributeValueMemberS{Value: "pendingInternal"},
			":d": &types.AttributeValueMemberS{Value: "pendingVenue"},
		},
	})

	log.DebugCtx(ctx, "query result for orders", out, err)

	if err != nil {
		log.WarnfCtx(ctx, "error listing orders for %s - %v", userId, err)
		return nil, fmt.Errorf("dynamodb could not query: %w", err)
	}

	log.DebugCtx(ctx, &out.Items)
	err = attributevalue.UnmarshalListOfMaps(out.Items, &orders)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal items: %w", err)
	}
	log.DebugCtx(ctx, "unmarshalled orders", &orders)

	return orders, nil
}
