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
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/coinbase-samples/ib-api-go/log"
)

func (am *Middleware) MakeHttpHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// allow health checks
			if r.RequestURI == "/health" || strings.Contains(r.RequestURI, "/ws?alias=") {
				next.ServeHTTP(w, r)
				return
			}

			ctx := r.Context()
			// Check header for bearer auth: "Authorization: Bearer <access token>"
			authorization := r.Header.Get("Authorization")
			bearerToken := extractBearerToken(authorization)
			if len(bearerToken) == 0 {
				log.DebugCtx(ctx, "missing bearer token")
				unauthenticatedResponse(w)
				return
			}

			user, err := am.Cip.GetUser(context.Background(), &cognitoidentityprovider.GetUserInput{
				AccessToken: aws.String(bearerToken),
			})

			if err != nil {
				log.DebugfCtx(ctx, "invalid bearer token: %v", err)
				unauthenticatedResponse(w)
				return
			}
			r = r.WithContext(addUserToContext(ctx, user))
			next.ServeHTTP(w, r)
		})
	}
}

func unauthenticatedResponse(w http.ResponseWriter) {
	w.WriteHeader(401)
	w.Write([]byte("unauthenticated"))
}

func extractBearerToken(authorization string) string {
	if authorization == "" {
		return ""
	}
	parts := strings.Split(authorization, "Bearer")
	if len(parts) != 2 {
		parts = strings.Split(authorization, "bearer")
		if len(parts) != 2 {
			return ""
		}
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
