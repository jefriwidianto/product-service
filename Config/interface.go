package Config

import (
	"database/sql"
)

type ConfigSettingSql struct {
	Environment
}

type Db interface {
	InitDB()
}

// maping all connection DB sql
var SqlConnection *sql.DB

func (d DbSqlConfigName) Get() *sql.DB {
	return SqlConnection
}
