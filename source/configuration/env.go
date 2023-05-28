package configuration

import (
	"doYourLogin/internal/configuration"
	"time"
)

var (
	APPLICATION_ENVIRONMENT = configuration.MakeConfig("dev")

	//Database
	DATABASE_DRIVER   = configuration.MakeConfig("sqlserver")
	DATABASE_NAME     = configuration.MakeConfig("doyourlogin")
	DATABASE_USERNAME = configuration.MakeConfig("sa")
	DATABASE_PASSWORD = configuration.MakeConfig("DoYourLoginEuOdeioReact2x")
	DATABASE_SOURCE   = configuration.MakeConfig("sqlserver://sa:DoYourLoginEuOdeioReact2x@localhost:1433?database=doyourlogin")

	//Jwt
	JWT_SECRET_KEY   = configuration.MakeConfig("my-jwt-secret-key")
	JWT_IDENTITY_KEY = configuration.MakeConfig("name")
	JWT_REALM        = configuration.MakeConfig("realm")
	JWT_TIMEOUT      = configuration.MakeConfig(time.Hour.String())
	JWT_MAX_REFRESH  = configuration.MakeConfig(time.Hour.String())

	//Email
	SMTP     = configuration.MakeConfig("smtp.gmail.com")
	PORT     = configuration.MakeConfig("587")
	EMAIL    = configuration.MakeConfig("vrossado1@gmail.com")
	PASSWORD = configuration.MakeConfig("ktdfeinefggyfjrx")

	//Mongo
	MONGO_CONNECTION_STRING = configuration.MakeConfig("mongodb://localhost:27017")
	MONGO_DATABASE_NAME     = configuration.MakeConfig("Logs")
	MONGO_COLLECTION_NAME   = configuration.MakeConfig("Logs")
	MONGO_USER              = configuration.MakeConfig("root")
	MONGO_PASSWORD          = configuration.MakeConfig("SeDerErroEscreveAwait")
)
