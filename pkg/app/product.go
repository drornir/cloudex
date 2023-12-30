package app

import (
	"context"
	"fmt"

	"github.com/drornir/cloudex/pkg/db"
	"github.com/drornir/cloudex/pkg/product"
)

func (a *App) BuyProduct(ctx context.Context, productName string) (int64, error) {
	errorW := func(err error) error {
		return fmt.Errorf("buying product %q: %w", productName, err)
	}

	user, err := UserFromContext(ctx)
	if err != nil {
		return 0, errorW(err)
	}

	prod, err := a.getProductByName(ctx, productName)
	if err != nil {
		return 0, errorW((err))
	}

	license, err := prod.NewLicense(ctx)
	if err != nil {
		return 0, errorW(fmt.Errorf("creating new license: %w", err))
	}

	id, err := a.assignLicenseToUser(ctx, user, license)
	if err != nil {
		return 0, errorW(err)
	}

	return id, nil
}

func (a *App) getProductByName(ctx context.Context, productName string) (product.Product, error) {
	prod, ok := product.Collection()[productName]
	if !ok {
		return nil, fmt.Errorf("product %q doesn't exist", productName)
	}

	return prod, nil
}

func (a *App) assignLicenseToUser(ctx context.Context, user User, license product.License) (int64, error) {
	id, err := a.DB.InsertLicense(ctx, db.InsertLicenseParams{
		Product:     license.Product().Name(),
		User:        user.ID,
		Credentials: license.Credentials(),
	})
	if err != nil {
		return 0, ErrorSQL{Msg: "InsertLicense", Err: err}
	}
	return id, nil
}
