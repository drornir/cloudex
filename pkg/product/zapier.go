package product

import "context"

type Zapier struct {
}

func (Zapier) Name() string {
	return "Zapier"
}

func (Zapier) Description() string {
	return "Zapier, you know"
}

func (Zapier) NewLicense(ctx context.Context) (License, error) {
	panic("unim")
}
