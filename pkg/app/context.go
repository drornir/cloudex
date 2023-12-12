package app

import "context"

type ctxKeyType string

const userCtxKey = "user"

type User struct {
	ID string
}

func UserFromContext(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userCtxKey).(User)
	if !ok {
		return u, ErrorUserIsMissing{Context: ctx}
	}
	return u, nil
}

func ContextWithUser(ctx context.Context, u User) context.Context {
	return context.WithValue(ctx, userCtxKey, u)
}
