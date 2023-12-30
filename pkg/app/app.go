package app

import "github.com/drornir/cloudex/pkg/db"

type App struct {
	DB *db.DB

	Vault *HashicorpVault
}
