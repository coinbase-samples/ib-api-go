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

package auth

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/coinbase-samples/ib-api-go/model"
)

func addUserToContext(ctx context.Context, user *cognitoidentityprovider.GetUserOutput) context.Context {
	var authedUser = model.User{}
	for _, attr := range user.UserAttributes {
		if *attr.Name == "sub" {
			authedUser.Id = *attr.Value
		} else if *attr.Name == "email" {
			authedUser.Email = *attr.Value
		}
	}
	log.DebugfCtx(ctx, "adding user to context: %v", authedUser)
	return context.WithValue(ctx, model.UserCtxKey, authedUser)
}
