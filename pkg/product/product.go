package product

import "context"

type Product interface {
	Name() string
	NewLicense(ctx context.Context) (License, error)
}

func Products() map[string]Product {
	return map[string]Product{
		Example{}.Name(): Example{},
		Zapier{}.Name():  Zapier{},
	}
}
