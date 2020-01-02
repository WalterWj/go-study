package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhost         = flag.String("h", "127.0.0.1", "IP for DB")
	dbusername     = flag.String("u", "root", "DB User")
	dbpassword     = flag.String("p", "", "Password")
	dbport         = flag.Int("P", 4000, "DB Port")
	dbstatus       = flag.Int("s", 10080, "DB Status Port")
	dbname         = flag.String("d", "test", "DB name")
	tbname         = flag.String("t", "test", "Table name")
	driver         = "mysql"
	dataSourceName = ""
)

func init() {
	flag.Parse()
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", *dbusername, *dbpassword, *dbhost, *dbport, *dbname)
}

func main() {
	// parser flag
	flag.Parse()
	start := time.Now()

	db := getdb()

	// rows, err := db.Query(fmt.Sprintf("show create table %s", *tbname))
	rows, err := db.Query("show tables")

	if err != nil {
		fmt.Println("Select err: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var content string
		rows.Scan(&content)
		// println(content)
		// parserTable(content)
	}
	// parserTable("test")
	// writeFile("tmp", "test.txt", "test", "nomal")
	// writeFile("tmp", "test.txt", "注释", "")
	// parserVersion()
	tmp := parserState(*dbname, *tbname, *dbhost, *dbstatus)
	println(tmp)
	costGetTime := time.Since(start)
	fmt.Printf("get values time is %s \n", costGetTime)
}

func parserState(dbname string, tbname string, dbhost string, dbstatus int) string {
	var pdURL string
	pdURL = fmt.Sprintf("http://%s:%d/stats/dump/%s/%s", dbhost, dbstatus, dbname, tbname)
	ret, err := http.Get(pdURL)
	if err != nil {
		panic(err)
	}
	defer ret.Body.Close()

	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(reflect.TypeOf(body))
	return string(body)
}

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
		f.WriteString(content + "\n")
	} else {
		content = "/* " + content + "*/"
		f.WriteString(content + "\n")
	}
}

// create table
func parserTable(tableName string) string {
	// parser flag
	flag.Parse()

	db := getdb()
	rows, err := db.Query(fmt.Sprintf("show create table %s", tableName))
	if err != nil {
		fmt.Println("Select err: ", err.Error())
	}
	defer rows.Close()

	var _table, content string
	for rows.Next() {
		rows.Scan(&_table, &content)
		// println(_table, content)
	}

	return content
}

// get version
func parserVersion() string {
	// parser flag
	flag.Parse()

	db := getdb()
	rows, err := db.Query("select tidb_version()")
	if err != nil {
		fmt.Println("Get TiDB version err: ", err.Error())
	}
	defer rows.Close()

	var content string
	for rows.Next() {
		rows.Scan(&content)
		// println(content)
	}

	return content
}

// connect DB
func getdb() *sql.DB {
	// get sql.DB
	db, err := sql.Open(driver, dataSourceName)

	if err != nil {
		fmt.Println("Connect err: ", err.Error())
	}

	return db
}
