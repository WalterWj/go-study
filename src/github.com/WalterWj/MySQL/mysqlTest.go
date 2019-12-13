package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhost     = "127.0.0.1:4000"
	dbusername = "root"
	dbpassword = ""
	dbname     = "test"
)

func main() {
	Insert("chain", "dev", "1")
	Insert("chain", "dev", "2")
	Insert("iris", "test", "1")
	Insert("iris", "test", "2")
}

func GetDB() *sql.DB {
	// get sql.DB
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbusername, dbpassword, dbhost, dbname))
	// CheckErr(err)
	println(err)
	return db
}

func Insert(username, departname, method string) bool {
	// insert data
	db := GetDB()
	defer db.Close()

	if method == "1" {
		_, err := db.Exec("insert into test2(name,dev,id) values(?,?,?)", username, departname, '1')
		if err != nil {
			fmt.Println("insert err: ", err.Error())
			return false
		}
		fmt.Println("insert success!")
		return true
	} else if method == "2" {
		stmt, err := db.Prepare("INSERT test2 SET name=?,dev=?,id=?")
		if err != nil {
			fmt.Println("insert prepare error: ", err.Error())
			return false
		}
		_, err = stmt.Exec(username, departname, 2)
		if err != nil {
			fmt.Println("insert exec error: ", err.Error())
			return false
		}
		fmt.Println("insert success!")
		return true
	}
	return false
}
