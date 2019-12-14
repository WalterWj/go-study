package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhost         = flag.String("h", "127.0.0.1", "IP for DB")
	dbusername     = flag.String("u", "root", "DB User")
	dbpassword     = flag.String("p", "", "Password")
	dbport         = flag.Int("P", 4000, "DB Port")
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

	rows, err := db.Query(fmt.Sprintf("SELECT * from %s", *tbname))

	if err != nil {
		fmt.Println("Select err: ", err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	//values：一行的所有值,把每一行的各个字段放到values中，values长度==列数
	values := make([]sql.RawBytes, len(columns))
	// print(len(values))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//存所有行的内容totalValues
	totalValues := make([][]string, 0)
	defer db.Close()
	for rows.Next() {

		//存每一行的内容
		var s []string

		//把每行的内容添加到scanArgs，也添加到了values
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for _, v := range values {
			s = append(s, string(v))
			// print(len(s))
		}
		totalValues = append(totalValues, s)
	}
	costGetTime := time.Since(start)
	fmt.Printf("get values time is %s \n", costGetTime)

	// writeToCSV(*tbname+".csv", columns, totalValues)

	costTotal := time.Since(start)
	fmt.Printf("write csv cost time is %s", costTotal-costGetTime)
}

// connect DB
func getdb() *sql.DB {
	// get sql.DB
	db, err := sql.Open(driver, dataSourceName)
	// CheckErr(err)
	if err != nil {
		fmt.Println("Connect err: ", err.Error())
	}
	return db
}

//writeToCSV
func writeToCSV(file string, columns []string, totalValues [][]string) {
	f, err := os.Create(file)
	// fmt.Println(columns)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	//f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	for i, row := range totalValues {
		//第一次写列名+第一行数据
		if i == 0 {
			w.Write(columns)
			w.Write(row)
		} else {
			w.Write(row)
		}
	}
	w.Flush()
	fmt.Println("处理完毕：", file)
}
