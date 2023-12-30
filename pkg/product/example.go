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

func (e Example) NewLicense(ctx context.Context) (License, error) {
	b := make([]byte, 32)
	rand.Read(b)
	credentials := base64.RawURLEncoding.EncodeToString(b)
	return ExampleLicense{
		credentials: credentials,
	}, nil
}

type ExampleLicense struct {
	credentials string
}

func (ExampleLicense) Product() Product {
	return Example{}
}

func (e ExampleLicense) Credentials() string {
	return e.credentials
}
