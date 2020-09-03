package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var (
	dbhost, dbname, dbusername, dbpassword string
	dbport                                 int
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
		for _, tableName := range res {
			showQ := fmt.Sprintf("show create table %s", tableName)
			rows, err := db.Query(showQ)
			ifErrWithLog(err)
			for rows.Next() {
				var t, Ct string
				err := rows.Scan(&t, &Ct)
				ifErrWithLog(err)
				fmt.Printf("%s;\n", Ct)
				d1 := []byte(Ct)
				err = ioutil.WriteFile(fmt.Sprintf("%s-%s.sql", dbname, tableName), d1, 0644)
				ifErrWithLog(err)
			}
			rows.Close()
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

// Get table schema information
func getTables(db *sql.DB) map[int]string {
	var r = make(map[int]string)
	rows, err := db.Query(tablesQ)
	ifErrWithLog(err)
	defer rows.Close()
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		ifErrWithLog(err)
		r[n] = t
		n++
	}
	return r
}
