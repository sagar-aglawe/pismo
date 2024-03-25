package providers

import (
	"self-projects/pismo/pkg/database/postgres"
	"sync"

	"self-projects/pismo/configs"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const PostgresClient = "postgres_client"

type PostgresProvider struct {
	Db        *gorm.DB
	Initiated bool
	mutex     *sync.Mutex
}

func init() {
	GlobalRegistry.Set(PostgresClient, &PostgresProvider{Initiated: false, mutex: &sync.Mutex{}})
}

func (pp *PostgresProvider) New(c *ProviderConfig) {

	pp.mutex.Lock()
	defer pp.mutex.Unlock()

	postgresProvider := GlobalRegistry.Get(PostgresClient)
	if !postgresProvider.IsInitiated() {
		gormDb := postgres.New(&c.PostgresClientConfig)
		postgresProvider := &PostgresProvider{Db: gormDb, Initiated: true}
		GlobalRegistry.Set(PostgresClient, postgresProvider)
	}
}

func (pp *PostgresProvider) IsInitiated() bool {
	return pp.Initiated
}

func getPostgresConfig() postgres.Config {
	return postgres.Config{
		Host:     viper.GetString(configs.PostgresHost),
		UserName: viper.GetString(configs.PostgresUser),
		PassWord: viper.GetString(configs.PostgresPassword),
		DbName:   viper.GetString(configs.PostgresDbName),
		Port:     viper.GetString(configs.PostgresPort),
	}
}
