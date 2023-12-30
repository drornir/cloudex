package app

import (
	"context"
	"fmt"
)

type Product interface {
	Name() string
}

type License interface {
	Product() Product
	Credentials() string
}

type ZapierProduct struct {
}

func (ZapierProduct) Name() string {
	return "Zapier"
}

func (a *App) BuyProduct(ctx context.Context, productName string) error {
	errorW := func(err error) error {
		return fmt.Errorf("buying product %q: %w", productName, err)
	}

	user, err := UserFromContext(ctx)
	if err != nil {
		return errorW(err)
	}

	product, err := a.getProductByName(ctx, productName)
	if err != nil {
		return errorW((err))
	}

	license, err := a.newLicenseForProduct(ctx, product)
	if err != nil {
		return errorW(err)
	}

	err = a.assignLicenseToUser(ctx, user, license)
	if err != nil {
		return errorW(err)
	}

	return nil
}
func (a *App) getProductByName(ctx context.Context, productName string) (Product, error) {
	panic("unim")
}

func (a *App) newLicenseForProduct(ctx context.Context, product Product) (License, error) {
	panic("unim")
}

func (a *App) assignLicenseToUser(ctx context.Context, user User, license License) error {
	panic("unim")
}
