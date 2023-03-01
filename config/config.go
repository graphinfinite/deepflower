package config

import "github.com/spf13/viper"

type Configuration struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host,
	Port,
	User,
	Password,
	Dbname,
	Sslmode string
}

func Init() (Configuration, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, err
	}

	pgconf := PostgresConfig{Host: viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		User:     viper.GetString("postgres.user"),
		Password: viper.GetString("postgres.password"),
		Dbname:   viper.GetString("postgres.dbname")}

	return Configuration{Postgres: pgconf}, nil
}
