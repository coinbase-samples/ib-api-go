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

func (m *Repository) ListAssets(ctx context.Context, requestUserId string) ([]model.Asset, error) {
	var assets []model.Asset

	l := ctxlogrus.Extract(ctx)
	l.Debugln("fetching assets for user ", requestUserId)

	limitAmount := int32(6)
	out, err := m.Svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(m.App.AssetTableName),
		Limit:     &limitAmount,
	})

	l.Debugln("query result for assets", out, err)

	if err != nil {
		return assets, err
	}

	l.Debugln(&out.Items)
	err = attributevalue.UnmarshalListOfMaps(out.Items, &assets)
	if err != nil {
		return assets, err
	}

	return assets, nil
}

func (m *Repository) GetAsset(ctx context.Context, requestUserId, assetId string) (model.Asset, error) {
	var asset model.Asset

	l := ctxlogrus.Extract(ctx)
	l.Debugln("fetching assets for user ", requestUserId)

	out, err := m.Svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(m.App.AssetTableName),
		KeyConditionExpression: aws.String("assetId = :a"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":a": &types.AttributeValueMemberS{Value: assetId},
		},
	})

	if err != nil {
		return asset, err
	}

	l.Debugln("order by id found", &out.Items[0])
	err = attributevalue.UnmarshalMap(out.Items[0], &asset)
	if err != nil {
		return asset, err
	}
	l.Debugln("unmarshalled order by id", &asset)

	return asset, nil
}
