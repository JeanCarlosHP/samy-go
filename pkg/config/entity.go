package config

import "github.com/jeancarloshp/samy-go/pkg/database"

type Config struct {
	Enviroment string
	AppPort    string
	Db         database.Db
}
