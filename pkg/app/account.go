package app

import "context"

// This is the thing we sell - an account to zapier, GCP etc.
type Account interface {
	Product() *Product
	AccountID(context.Context) (string, error)
	NewAPIKey(ctx context.Context, name string) (string, error)
}
