package app

import (
	"context"
	"log"
)

type ctxKeyType string

const userCtxKey = "user"

type User struct {
	ID    int64
	Email string
}

func UserFromContext(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userCtxKey).(User)
	if !ok {
		return u, ErrorUserIsMissing{Context: ctx}
	}
	return u, nil
}

func ContextWithUser(ctx context.Context, u User) context.Context {
	log.Printf("Adding user to context: %#v", u)
	return context.WithValue(ctx, userCtxKey, u)
}
