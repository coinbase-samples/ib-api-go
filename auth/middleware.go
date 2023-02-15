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

package auth

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/coinbase-samples/ib-api-go/log"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

type AuthClient interface {
	GetUser(ctx context.Context, params *cognitoidentityprovider.GetUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserOutput, error)
}

type Middleware struct {
	Cip AuthClient
}

func (am *Middleware) InterceptorNew() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// allow health checks to pass through
		if info.FullMethod == "/grpc.health.v1.Health/Check" || info.FullMethod == "/grpc.health.v1.Health/Watch" {
			return handler(ctx, req)
		}

		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		log.DebugfCtx(ctx, "checking token, method: %s", info.FullMethod)
		if err != nil {
			log.Debugf("no bearer token in metadata - %v", err)
			return nil, fmt.Errorf("could not find bearer token from metadata: %w", err)
		}

		user, err := am.Cip.GetUser(ctx, &cognitoidentityprovider.GetUserInput{
			AccessToken: aws.String(token),
		})
		if err != nil {
			log.DebugfCtx(ctx, "could not validate token - %v", err)
			return nil, fmt.Errorf("could not validate token: %w", err)
		}

		ctx = addUserToContext(ctx, user)
		return handler(ctx, req)
	}
}
