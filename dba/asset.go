package dba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/coinbase-samples/ib-api-go/model"
)

func (m *Repository) ListAssets(ctx context.Context, requestUserId string) ([]model.Asset, error) {
	var assets []model.Asset

	log.CtxDebug(ctx, "fetching assets for user ", requestUserId)

	limitAmount := int32(6)
	out, err := m.Svc.Scan(context.Background(), &dynamodb.ScanInput{
		TableName: aws.String(m.App.AssetTableName),
		Limit:     &limitAmount,
	})

	log.CtxDebugf(ctx, "query result for assets: %v", out)

	if err != nil {
		log.CtxDebugf(ctx, "error listing assets - %v", err)
		return assets, err
	}

	log.CtxDebugf(ctx, "unmarshalled assets: %v", &out.Items)
	err = attributevalue.UnmarshalListOfMaps(out.Items, &assets)
	if err != nil {
		return assets, err
	}
	log.CtxDebugf(ctx, "returning final assets: %v", &assets)
	return assets, nil
}

func (m *Repository) GetAsset(ctx context.Context, requestUserId, assetId string) (model.Asset, error) {
	var asset model.Asset

	log.CtxDebug(ctx, "fetching assets for user ", requestUserId)

	out, err := m.Svc.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(m.App.AssetTableName),
		KeyConditionExpression: aws.String("assetId = :a"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":a": &types.AttributeValueMemberS{Value: assetId},
		},
	})

	if err != nil {
		return asset, err
	}

	log.CtxDebugf(ctx, "order by id found: %v", &out.Items[0])
	err = attributevalue.UnmarshalMap(out.Items[0], &asset)
	if err != nil {
		return asset, err
	}
	log.CtxDebugf(ctx, "returning final asset: %v", &asset)

	return asset, nil
}
