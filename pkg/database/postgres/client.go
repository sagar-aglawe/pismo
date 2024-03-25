package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Host     string `json:"host"`
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
	DbName   string `json:"db_name"`
	Port     string `json:"port"`
}

func New(c *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.Host,
		c.UserName,
		c.PassWord,
		c.DbName,
		c.Port,
	)

	postgresDb, err := gorm.Open(

		postgres.New(
			postgres.Config{
				DSN: dsn,
			},
		), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
		},
	)

	if err != nil {
		panic(fmt.Sprintf("not able to connect to database , error is : %s", err.Error()))
	}

	_, err = postgresDb.DB()
	if err != nil {
		panic(fmt.Sprintf("not able to connect to database , error is : %s", err.Error()))
	}

	return postgresDb
}
