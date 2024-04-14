package utils

import (
	"log"

	"github.com/spf13/viper"
)

type PostgresConfig struct {
	DB_USER string `mapstructure:"DB_USER"`
	DB_HOST string `mapstructure:"DB_HOST"`
	DB_PASS string `mapstructure:"DB_PASS"`
	DB_NAME string `mapstructure:"DB_NAME"`
	DB_PORT int    `mapstructure:"DB_PORT"`
}

func LoadPostgresConfig(path string) *PostgresConfig {
	viper.AddConfigPath(path)
	viper.SetConfigName("postgres")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	postgresConfig := &PostgresConfig{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to read postgres configuration: ", err)
	}
	if err := viper.Unmarshal(postgresConfig); err != nil {
		log.Fatal("failed to unmarshal postgres configuration: ", err)
	}

	return postgresConfig
}
