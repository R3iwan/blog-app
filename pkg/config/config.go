package config

import "github.com/spf13/viper"

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Config struct {
	Postgres   PostgresConfig
	Port       string
	JWT_Secret string
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Config{
		Postgres: PostgresConfig{
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetInt("POSTGRES_PORT"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			DBName:   viper.GetString("POSTGRES_DB"),
		},
		Port:       viper.GetString("PORT"),
		JWT_Secret: viper.GetString("JWT_SECRET"),
	}, nil

}
