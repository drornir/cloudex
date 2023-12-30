package product

import "context"

type Product interface {
	Name() string
	NewLicense(ctx context.Context) (License, error)
}

type License interface {
	Product() Product
	Credentials() string
}

func Collection() map[string]Product {
	return map[string]Product{
		Example{}.Name(): Example{},
		Zapier{}.Name():  Zapier{},
	}
}
