package Config

import (
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func (c ConfigSettingSql) InitDB() {
	//convert data object Config to buffer string
	var buffer bytes.Buffer
	buffer.WriteString(c.Databases.Username + ":" + c.Databases.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(c.Databases.Host + ":" + c.Databases.Port + ")/")
	buffer.WriteString("?charset=utf8")
	connection_string := buffer.String()

	Connection, err := sql.Open(c.Databases.Engine, connection_string)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//set max connection into databases
	Connection.SetMaxOpenConns(c.Databases.Maximum_connection)
	err = Connection.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
		return
	}
	log.Println("ENGINE " + c.Databases.Engine + " START....")

	//create connection if connection ping to databases not error
	SqlConnection = Connection

	//run migration Table
	err = migrate(SqlConnection)
	if err != nil {
		panic(err.Error())
		return
	}
}

func migrate(db *sql.DB) (err error) {
	tx, err := db.Begin()
	var queryAll []string
	query, err := os.ReadFile(dirPathMigration())
	sqlQuery := string(query)
	queryAll = strings.Split(sqlQuery, ";")
	for _, v := range queryAll {
		_, err = tx.Exec(v)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return
}

func dirPathMigration() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "../"+PathMigration+"20240402database_migration.sql")
}
