package model

type UserCtxKeyType string

const UserCtxKey UserCtxKeyType = "user"

type RequestIdCtxKeyType string

const RequestCtxKey RequestIdCtxKeyType = "requestId"

type User struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}
