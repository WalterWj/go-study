package command

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

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

// Write content
func writeFile(dirname string, fileName string, content string, mode string) {
	_, err := os.Stat(dirname)
	if err != nil {
		fmt.Printf("dir %s is not exist\n", dirname)
		// mkdir dir
		err := os.Mkdir(dirname, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir %s success!\n", dirname)
		}
	}

	f, err := os.OpenFile(path.Join(dirname, fileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("os OpenFile error: ", err)
		return
	}
	defer f.Close()
	if mode == "nomal" {
		f.WriteString(content + ";\n")
	} else if mode == "json" {
		f.WriteString(content + "\n")
	} else {
		content = "/*\n" + content + "\n*/"
		f.WriteString(content + "\n")
	}
}
