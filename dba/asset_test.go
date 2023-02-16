/**
 * Copyright 2022-present Coinbase Global, Inc.
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
	"testing"
)

func TestListAssets(t *testing.T) {
	repo := new(MockRepository)
	NewTestDBA(repo)

	resp, err := repo.ListAssets(context.Background(), ListAssetsFound)

	if err != nil {
		t.Fatal("unexpected error returned from function invocation")
	}

	if resp[0].Name != "Bitcoin" {
		t.Fatal("expected name")
	}

	if resp[0].Ticker != "BTC" {
		t.Fatal("expected userId to match")
	}
}

func TestListAssetsNotfound(t *testing.T) {
	repo := new(MockRepository)
	NewTestDBA(repo)

	resp, err := repo.ListAssets(context.Background(), ListAssetsNotFound)

	if err == nil {
		t.Fatal("expected error returned from function invocation")
	}

	if resp != nil {
		t.Fatal("expected no items returned")
	}
}

func TestGetAsset(t *testing.T) {
	repo := new(MockRepository)
	NewTestDBA(repo)

	resp, err := repo.GetAsset(context.Background(), GetAsset, "a")

	if err != nil {
		t.Fatal("unexpected error returned from function invocation")
	}

	if resp.Name != "Bitcoin" {
		t.Fatal("expected name")
	}

	if resp.AssetId != "a" {
		t.Fatal("expected assetId to match")
	}
}

func TestGetAssetsNotFound(t *testing.T) {
	repo := new(MockRepository)
	NewTestDBA(repo)

	_, err := repo.GetAsset(context.Background(), GetAssetNotFound, "a")

	if err == nil {
		t.Fatal("expected error returned from function invocation")
	}

}
