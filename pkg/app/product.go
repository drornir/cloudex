package app

import (
	"context"
	"fmt"

	"github.com/drornir/cloudex/pkg/db"
	"github.com/drornir/cloudex/pkg/product"
)

func (a *App) BuyProduct(ctx context.Context, productName string) (product.LicenseAndMeta, error) {
	errorW := func(err error) error {
		return fmt.Errorf("buying product %q: %w", productName, err)
	}

	user, err := UserFromContext(ctx)
	if err != nil {
		return product.LicenseAndMeta{}, errorW(err)
	}

	prod, err := a.getProductByName(ctx, productName)
	if err != nil {
		return product.LicenseAndMeta{}, errorW((err))
	}

	license, err := prod.NewLicense(ctx)
	if err != nil {
		return product.LicenseAndMeta{}, errorW(fmt.Errorf("creating new license: %w", err))
	}

	l, err := a.assignLicenseToUser(ctx, user, license)
	if err != nil {
		return product.LicenseAndMeta{}, errorW(err)
	}

	return l, nil
}

func (a *App) getProductByName(ctx context.Context, productName string) (product.Product, error) {
	prod, ok := product.Products()[productName]
	if !ok {
		return nil, fmt.Errorf("product %q doesn't exist", productName)
	}

	return prod, nil
}

func (a *App) assignLicenseToUser(ctx context.Context, user User, license product.License) (product.LicenseAndMeta, error) {
	dbl, err := a.DB.InsertLicense(ctx, db.InsertLicenseParams{
		Product:     license.Product().Name(),
		User:        user.ID,
		Credentials: license.Credentials(),
	})
	if err != nil {
		return product.LicenseAndMeta{},
			ErrorSQL{Msg: "InsertLicense", Err: err}
	}
	l := UnmarshalLicense(dbl)
	return l, nil
}

func UnmarshalLicense(l db.Licenses) product.LicenseAndMeta {
	prod := product.Products()[l.Product]
	switch prod.(type) {
	case product.Example:
		return product.LicenseAndMeta{
			License: product.ExampleLicense{
				Creds: l.Credentials,
			},
			Meta: product.LicenseMeta{
				ID:     l.ID,
				UserID: l.User,
			},
		}
	// case product.Zapier:
	default:
		panic(fmt.Errorf("unsupported prod type %q in db.License id=%d", l.Product, l.ID))
	}
}
