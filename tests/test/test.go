package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhost     = flag.String("h", "127.0.0.1", "IP for DB")
	dbusername = flag.String("u", "root", "DB User")
	dbpassword = flag.String("p", "", "Password")
	dbport     = flag.Int("P", 4000, "DB Port")
	dbstatus   = flag.Int("s", 10080, "DB Status Port")
	dbname     = flag.String("d", "test", "DB name")
	tbname     = flag.String("t", "test", "Table name")
)

var (
	driver         = "mysql"
	dataSourceName = ""
	dirname        = time.Now().Unix()
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

	}

	parserVersion()
	parserTable(*tbname)
	parserState(*dbname, *tbname)

	// if *tbname != "test" {
	// 	println("1")
	// } else {
	// 	println("2")
	// }
	costGetTime := time.Since(start)
	fmt.Printf("Cost time is %s \n", costGetTime)
	// a := strings.Split("1;2;3;4", ";")
	// for _, _a := range a {
	// 	println(_a)
	// }
	// println(len(*dbname))
	// println(dirname)
}

func parserState(dbname string, tbname string) string {
	var pdURL string
	pdURL = fmt.Sprintf("http://%s:%d/stats/dump/%s/%s", *dbhost, *dbstatus, dbname, tbname)
	ret, err := http.Get(pdURL)
	if err != nil {
		panic(err)
	}
	defer ret.Body.Close()

	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%s.%s.json", dbname, tbname)
	// fmt.Println(reflect.TypeOf(body))
	_dirname := "stats-" + strconv.FormatInt(dirname, 10) + "/stats"
	_content := string(body)
	writeFile(_dirname, fileName, _content, "json")

	_dirname = "stats-" + strconv.FormatInt(dirname, 10)
	_content = fmt.Sprintf("LOAD STATS 'stats/%s'", fileName)
	writeFile(_dirname, "schema.sql", _content, "nomal")

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
		f.WriteString(content + ";\n")
	} else if mode == "json" {
		f.WriteString(content + "\n")
	} else {
		content = "/*\n" + content + "\n*/"
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
	_dirname := "stats-" + strconv.FormatInt(dirname, 10)
	writeFile(_dirname, "schema.sql", tableName, "")
	writeFile(_dirname, "schema.sql", content, "nomal")

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
	_dirname := "stats-" + strconv.FormatInt(dirname, 10)
	writeFile(_dirname, "schema.sql", content, "")

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
