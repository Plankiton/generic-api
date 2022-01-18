package env

import "os"

// SQL_URL defines the mysql host to connection
var SQL_PORT, SQL_USER, SQL_PASS, SQL_DB, SQL_URL, SQL_HOST string
var SQL_USE_SQLITE bool

func init() {
	SQL_USE_SQLITE = os.Getenv("SQL_USE_SQLITE") == "1"

	SQL_PORT = os.Getenv("SQL_PORT")
	SQL_HOST = os.Getenv("SQL_HOST")
	SQL_USER = os.Getenv("SQL_USER")
	SQL_PASS = os.Getenv("SQL_PASS")
	SQL_DB = os.Getenv("SQL_DB")
	SQL_URL = os.Getenv("SQL_URL")
}
