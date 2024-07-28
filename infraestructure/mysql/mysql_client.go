package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chava.gnolasco/polaris/infraestructure/env"
	"github.com/chava.gnolasco/polaris/infraestructure/log"
	_ "github.com/go-sql-driver/mysql"
)

// MySQLClient is the MySQL client.
var mysqlClient *sql.DB

// GetMySQLClient returns the MySQL client.
func GetMySQLClient() *sql.DB {
	return mysqlClient
}

// Creates a new MySQL client.
func init() {
	log.Info("[Mysql] Initializing the connection...")
	env := env.GetEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		env.MYSQL_USER, env.MYSQL_PASSWORD, env.MYSQL_URL, env.MYSQL_DB_NAME)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Error("Error trying to connect to database: " + err.Error())
		os.Exit(1)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(0) // No tiempo máximo para la conexión

	if err := db.Ping(); err != nil {
		log.Error("Error on database ping: " + err.Error())
		os.Exit(1)
	}

	mysqlClient = db
}
