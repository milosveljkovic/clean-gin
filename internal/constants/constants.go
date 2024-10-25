package constants

import "github.com/gin-gonic/gin"

const (
	CleanGinDefaultPort = "8123"
	CleanGinDefaultMode = gin.DebugMode
)

// database
const (
	PostgreqlDBPort      = "POSTGRESQL_DB_PORT"
	PostgresqlDBHost     = "POSTGRESQL_DB_HOST"
	PostgresqlDBName     = "POSTGRESQL_DB_NAME"
	PostgresqlDBPassword = "POSTGRESQL_DB_PASSWORD"
	PostgresqlDBUser     = "POSTGRESQL_DB_USER"
	PostgresqlDBSslMode  = "POSTGRESQL_DB_SSL_MODE"
)
