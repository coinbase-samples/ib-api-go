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
		log.CtxDebugf(ctx, "user attr: %s - %s", *attr.Name, *attr.Value)
		if *attr.Name == "sub" {
			authedUser.Id = *attr.Value
		} else if *attr.Name == "email" {
			authedUser.Email = *attr.Value
		}
	}
	log.CtxDebugf(ctx, "adding user to context: %v", authedUser)
	return context.WithValue(ctx, model.UserCtxKey, authedUser)
}
