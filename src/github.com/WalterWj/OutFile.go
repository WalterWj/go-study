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
	filter         = flag.String("f", "", "filter condition")
	clumn          = flag.String("c", "*", "Columns to be exported")
	driver         = "mysql"
	dataSourceName = ""
)

func init() {
	flag.Parse()
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", *dbusername, *dbpassword, *dbhost, *dbport, *dbname)
}

func main() {
	start := time.Now()
	// cmd := fmt.Sprintf("select * from %s", *tbname)
	cmd := fmt.Sprintf("select %s from %s %s", *clumn, *tbname, *filter)

	db, err := sql.Open(driver, dataSourceName)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	outFile(db, cmd, *tbname+".csv")
	costTime := time.Since(start)
	fmt.Printf("Thread cost time is %s \n", costTime)
}

func outFile(db *sql.DB, cmd string, fileNmae string) {
	start := time.Now()

	rows, err := db.Query(cmd)

	if err != nil {
		fmt.Println("Select err: ", err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	//values：一行的所有值,把每一行的各个字段放到values中，values长度==列数
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//存所有行的内容totalValues
	totalValues := make([][]string, 0)

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
		}
		totalValues = append(totalValues, s)
	}

	costGetTime := time.Since(start)
	fmt.Printf("get values time is %s \n", costGetTime)

	writeToCSV(fileNmae, columns, totalValues)

	costTotal := time.Since(start)
	fmt.Printf("write csv cost time is %s \n", costTotal-costGetTime)
}

//writeToCSV
func writeToCSV(file string, columns []string, totalValues [][]string) {
	f, err := os.Create(file)

	defer f.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("开始写入内容：", file)
	w := csv.NewWriter(f)
	for i, row := range totalValues {
		// 第一次写列名+第一行数据
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
