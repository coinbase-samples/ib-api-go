/**
 * Copyright 2022 - Present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &orders); err != nil {
		return nil, fmt.Errorf("could not unmarshal items: %w", err)
	}
	log.DebugCtx(ctx, "unmarshalled orders", &orders)

	return orders, nil
}
