package command

import (
	"database/sql"
	"fmt"
	"log"

	// import mysql
	_ "github.com/go-sql-driver/mysql"
)

var id int

// Show : mysqlconnect
func Show(dsn string) {
	// dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", dbusername, dbpassword, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", dsn)
	err = db.QueryRow("select id from t where id = ?", 1).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	defer db.Close()
}
