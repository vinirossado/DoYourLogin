package configuration

import (
	"abrigos/internal/configuration"
	"time"
)

var (
	APPLICATION_ENVIRONMENT = configuration.MakeConfig("dev")

	//Database
	DATABASE_DRIVER   = configuration.MakeConfig("sqlserver")
	DATABASE_NAME     = configuration.MakeConfig("abrigo")
	DATABASE_USERNAME = configuration.MakeConfig("sa")
	DATABASE_PASSWORD = configuration.MakeConfig("root")
	DATABASE_SOURCE   = configuration.MakeConfig("sqlserver://sa:root@localhost:1433?database=abrigo")

	//Jwt
	JWT_SECRET_KEY   = configuration.MakeConfig("my-jwt-secret-key")
	JWT_IDENTITY_KEY = configuration.MakeConfig("name")
	JWT_REALM        = configuration.MakeConfig("realm")
	JWT_TIMEOUT      = configuration.MakeConfig(time.Hour.String())
	JWT_MAX_REFRESH  = configuration.MakeConfig(time.Hour.String())
)
