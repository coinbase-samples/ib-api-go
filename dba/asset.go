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

func (m *DynamoRepository) ListAssets(ctx context.Context, requestUserId string) ([]model.Asset, error) {
	var assets []model.Asset

	log.DebugCtx(ctx, "fetching assets for user ", requestUserId)

	limitAmount := int32(6)
	out, err := m.Svc.Scan(context.Background(), &dynamodb.ScanInput{
		TableName: aws.String(m.App.AssetTableName),
		Limit:     &limitAmount,
	})

	log.DebugfCtx(ctx, "query result for assets: %v", out)

	if err != nil {
		log.WarnfCtx(ctx, "error listing assets - %v", err)
		return nil, fmt.Errorf("dynamodb could not scan: %w", err)
	}

	log.DebugfCtx(ctx, "unmarshalled assets: %v", &out.Items)
	err = attributevalue.UnmarshalListOfMaps(out.Items, &assets)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal items: %w", err)
	}
	log.DebugfCtx(ctx, "returning final assets: %v", &assets)
	return assets, nil
}

func (m *DynamoRepository) GetAsset(ctx context.Context, requestUserId, assetId string) (model.Asset, error) {
	var asset model.Asset

	log.DebugCtx(ctx, "fetching assets for user ", requestUserId)

	out, err := m.Svc.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(m.App.AssetTableName),
		KeyConditionExpression: aws.String("assetId = :a"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":a": &types.AttributeValueMemberS{Value: assetId},
		},
	})

	if err != nil {
		log.WarnfCtx(ctx, "error fetching asset %s - %v", assetId, err)
		return asset, fmt.Errorf("dynamodb could not query: %w", err)
	}

	log.DebugfCtx(ctx, "order by id found: %v", &out.Items[0])
	err = attributevalue.UnmarshalMap(out.Items[0], &asset)
	if err != nil {
		return asset, fmt.Errorf("could not unmarshal item: %w", err)
	}
	log.DebugfCtx(ctx, "returning final asset: %v", &asset)

	return asset, nil
}
