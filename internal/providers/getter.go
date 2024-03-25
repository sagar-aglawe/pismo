package providers

import (
	"gorm.io/gorm"
)

func GetPostgresClient() *gorm.DB {
	provider := GlobalRegistry.GetOrResolve(PostgresClient)
	postgresClient := provider.(*PostgresProvider)
	return postgresClient.Db
}
