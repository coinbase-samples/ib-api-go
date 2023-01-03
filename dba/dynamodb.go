package dba

import (
	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/coinbase-samples/ib-api-go/config"
)

// Repo the repository used by dynamo
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	Svc *dynamodb.Client
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, cfg awsConfig.Config) *Repository {
	svc := setupService(a, cfg)

	return &Repository{
		App: a,
		Svc: svc,
	}
}

// NewDBA sets the repository for the handlers
func NewDBA(r *Repository) {
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
