package app

import (
	"io/fs"

	"github.com/drornir/cloudex/pkg/db"
)

type App struct {
	DB       *db.DB
	AssetsFS fs.FS

	Vault *HashicorpVault
}
