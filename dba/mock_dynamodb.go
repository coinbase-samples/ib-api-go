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
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/model"
)

var TestRepo *MockRepository

type MockRepository struct {
	App *config.AppConfig
	Svc *dynamodb.Client
}

func NewTestDBA(r *MockRepository) {
	TestRepo = r
}

var (
	ListAssetsNotFound = "4AC6E407-1D8E-4339-BA1C-862ACC58AC5E"
	ListAssetsFound    = "20032259-738B-40A7-AAD7-306B69AF88D4"
	GetAssetNotFound   = "E6096F2D-C706-42B6-B0E5-D7DD644ED079"
	GetAsset           = "AC032259-738B-40A7-AAD7-306B69AAB909"
	ListOrders         = "9165BBBF-02B8-408C-8345-04C34D8C1594"
	ListOrdersNotFound = "8FE47FE3-3B80-44EF-AEBF-DE9A1DF7491E"
)

func (m *MockRepository) ListAssets(ctx context.Context, requestUserId string) ([]model.Asset, error) {
	if requestUserId == ListAssetsNotFound {
		return nil, errors.New("error listing assets")
	}
	return []model.Asset{{Name: "Bitcoin", Ticker: "BTC"}}, nil
}

func (m *MockRepository) GetAsset(ctx context.Context, requestUserId, assetId string) (model.Asset, error) {
	if requestUserId == GetAssetNotFound {
		return model.Asset{}, errors.New("dynamodb could not query: ")
	}
	return model.Asset{Name: "Bitcoin", AssetId: assetId}, nil
}

func (m *MockRepository) ListOrders(ctx context.Context, userId string) ([]model.Order, error) {
	if userId == ListOrdersNotFound {
		return nil, errors.New("dynamodb could not query: ")
	}
	orders := []model.Order{{ProductId: "BTC-USD", Quantity: "1"}}
	return orders, nil
}
