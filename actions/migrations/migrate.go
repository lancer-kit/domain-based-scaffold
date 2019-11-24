package migrations

import (
	"github.com/lancer-kit/armory/db"
	migrate "github.com/rubenv/sql-migrate"
)

//go get -u github.com/lancer-kit/forge

//go:generate forge bindata --ignore .+\.go$ --pkg migrations -o bindata.go -i ./...
//go:generate gofmt -w bindata.go

func Migrate(connStr string, dir db.MigrateDir) (int, error) {
	db.SetAssets(migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "postgres",
	})
	return db.Migrate(connStr, dir)
}
