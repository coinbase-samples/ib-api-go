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
			if r.RequestURI == "/health" {
				next.ServeHTTP(w, r)
			}

			ctx := r.Context()
			// Check header for bearer auth: "Authorization: Bearer <access token>"
			authorization := r.Header.Get("Authorization")
			bearerToken := extractBearerToken(authorization)
			log.CtxDebugf(ctx, "got bearer token: %s", bearerToken)
			if len(bearerToken) == 0 {
				log.CtxDebug(ctx, "missing bearer token")
				unauthenticatedResponse(w)
				return
			}

			user, err := am.Cip.GetUser(context.Background(), &cognitoidentityprovider.GetUserInput{
				AccessToken: aws.String(bearerToken),
			})

			if err != nil {
				log.CtxDebugf(ctx, "invalid bearer token: %v", err)
				unauthenticatedResponse(w)
				return
			}
			log.CtxDebugf(ctx, "fetched cognito user: %v", user)
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
