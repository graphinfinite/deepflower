package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Db       DbConfig
	Auth     AuthConfig
	Server   ServerConfig
	Telegram TelegramConfig
}

type AuthConfig struct {
	Hash_salt   int
	Signing_key int
	Token_ttl   int
}

type DbConfig struct {
	Host,
	Port,
	User,
	Password,
	Dbname,
	Sslmode,
	Psql string
}

type ServerConfig struct {
	Host,
	Port string
}

type TelegramConfig struct {
	Token,
	Boturl string
}

func Init() (Configuration, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, err
	}

	viper.AutomaticEnv()

	dbconf := DbConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		User:     viper.GetString("postgres.user"),
		Password: viper.GetString("postgres.password"),
		Dbname:   viper.GetString("postgres.dbname")}

	dbconf.Psql = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbconf.Host,
		dbconf.Port,
		dbconf.User,
		dbconf.Password,
		dbconf.Dbname)

	authconf := AuthConfig{
		Hash_salt:   viper.GetInt("auth.hash_salt"),
		Signing_key: viper.GetInt("auth.signing_key"),
		Token_ttl:   viper.GetInt("auth.token_ttl")}

	srvconf := ServerConfig{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}
	tgconf := TelegramConfig{
		Token:  viper.GetString("telegram.token"),
		Boturl: viper.GetString("telegram.boturl"),
	}

	return Configuration{Db: dbconf, Auth: authconf, Server: srvconf, Telegram: tgconf}, nil
}
