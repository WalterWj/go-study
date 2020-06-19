package command

import (
	"database/sql"
	"fmt"
)

// mysqlconnect
func Show(dbusername string, dbhost string, dbpassword string, dbname string, dbport int) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", dbusername, dbpassword, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", dsn)
}
