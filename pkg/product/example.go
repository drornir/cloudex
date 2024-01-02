package product

import (
	"context"
	"crypto/rand"
	"encoding/base64"
)

type Example struct {
}

func (e Example) Name() string {
	return "Example"
}

func (e Example) Description() string {
	return "An example product to try things with."
}

func (e Example) NewLicense(ctx context.Context) (License, error) {
	b := make([]byte, 32)
	rand.Read(b)
	credentials := base64.RawURLEncoding.EncodeToString(b)
	return ExampleLicense{
		Creds: credentials,
	}, nil
}

type ExampleLicense struct {
	ID     int64
	UserID int64
	Creds  string
}

func (ExampleLicense) Product() Product {
	return Products()[Example{}.Name()]
}

func (e ExampleLicense) Credentials() string {
	return e.Creds
}
