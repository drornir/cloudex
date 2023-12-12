package app

import "context"

type App struct {
	Vault *HashicorpVault
}

func (a *App) RegisterToProduct(ctx context.Context, prod Product) (Account, error) {
	user, err := UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	_ = user
	panic("unimplemented")
}
