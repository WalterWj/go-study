package command

import (
	"database/sql"
	"fmt"
	"log"
)

var id int

// Show : mysqlconnect
func Show(dbusername string, dbhost string, dbpassword string, dbname string, dbport int) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", dbusername, dbpassword, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", dsn)
	err = db.QueryRow("select id from user where id = ?", 1).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
