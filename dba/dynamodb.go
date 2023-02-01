package dba

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/model"
)

type Repository interface {
	ListOrders(ctx context.Context, userId string) ([]model.Order, error)
	ListAssets(ctx context.Context, requestUserId string) ([]model.Asset, error)
	GetAsset(ctx context.Context, requestUserId, assetId string) (model.Asset, error)
}

// Repo the repository used by dynamo
var Repo *DynamoRepository

// Repository is the repository type
type DynamoRepository struct {
	App *config.AppConfig
	Svc *dynamodb.Client
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, cfg awsConfig.Config) *DynamoRepository {
	svc := setupService(a, cfg)

	return &DynamoRepository{
		App: a,
		Svc: svc,
	}
}

// NewDBA sets the repository for the handlers
func NewDBA(r *DynamoRepository) {
	Repo = r
}

func setupService(a *config.AppConfig, cfg awsConfig.Config) *dynamodb.Client {
	var svc *dynamodb.Client

	if a.IsLocalEnv() {
		svc = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.EndpointResolver = dynamodb.EndpointResolverFromURL(a.DatabaseEndpoint)
		})
	} else {
		svc = dynamodb.NewFromConfig(cfg)
	}

	return svc
}
