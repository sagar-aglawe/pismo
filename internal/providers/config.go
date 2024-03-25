package providers

import (
	"self-projects/pismo/pkg/database/postgres"
)

type ProviderConfig struct {
	PostgresClientConfig postgres.Config
}
