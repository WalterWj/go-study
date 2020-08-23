package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	dbhost, dbname, tbname, dbusername, dbpassword string
)

const (
	tablesQ = "show tables"
)

// statsDumpCmd represents the statsDump command
var statsDumpCmd = &cobra.Command{
	Use:   "statsDump",
	Short: "Export statistics and table structures",
	Long:  `Export statistics and table structures`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("statsDump called")
		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", dbusername, dbpassword, dbhost, dbport, dbname)
		db := mysqlConnect(dsn)
		res := getTables(db)
		for _, Tname := range res {
			fmt.Println(Tname)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsDumpCmd)
	statsDumpCmd.Flags().StringVarP(&dbusername, "dbusername", "u", "root", "Database user")
	statsDumpCmd.Flags().StringVarP(&dbname, "dbname", "d", "test", "Database name")
	statsDumpCmd.Flags().StringVarP(&dbhost, "dbhost", "H", "127.0.0.1", "Database host")
	statsDumpCmd.Flags().StringVarP(&dbpassword, "dbpassword", "p", "", "Database passowrd")
	statsDumpCmd.Flags().IntVarP(&dbport, "dbport", "P", 4000, "Database port")
}

func getTables(db *sql.DB) map[int]string {
	var r = make(map[int]string)
	rows, err := db.Query(tablesQ)
	ifErrWithLog(err)
	defer rows.Close()
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		n := len(t)
		ifErrWithLog(err)
		r[n] = t
	}
	err = rows.Err()
	ifErrWithLog(err)
	return r
}

func getTables(db *sql.DB)  map[string]string{
	var r = make(map[string]string)
	rows, err := db.Query(tables_q)
	ifErrWithLog(err)
	defer rows.Close()

	for rows.Next() {
		var n, t string
		err := rows.Scan(&t)
		ifErrWithLog(err)
		r[n] = t
	}
	err = rows.Err()
	ifErrWithLog(err)
	return r
}