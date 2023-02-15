# Copyright 2022 - Present Coinbase Global, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#  http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

REGION ?= us-east-1
PROFILE ?= sa-infra
ENV_NAME ?= dev

ACCOUNT_ID := $(shell aws sts get-caller-identity --profile $(PROFILE) --query 'Account' --output text)

.PHONY: docker-build
docker-build:
	@docker build --platform linux/amd64 --build-arg REGION=$(REGION) --build-arg ENV_NAME=$(ENV_NAME) --build-arg ACCOUNT_ID=$(ACCOUNT_ID) .

.PHONY: start
start:
	go run cmd/server/*.go

.PHONY: docker-build-local
docker-build-local:
	@docker build --tag ib-api-go:local --build-arg REGION=$(REGION) --build-arg ENV_NAME=local --build-arg ACCOUNT_ID=$(ACCOUNT_ID) .

.PHONY: docker-start-local
docker-start-local:
	@docker run --net ib-system_default -p 8443:8443 --env-file .env -d ib-api-go:local
