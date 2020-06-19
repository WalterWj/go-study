package cmd

import (
	"fmt"

	"github.com/WalterWj/go-study/command"
	"github.com/spf13/cobra"
)

var (
	dbhost, dbname, tbname, dbusername, dbpassword, id string
	dbport, dbstatus                                   int
)

// statsDumpCmd represents the statsDump command
var statsDumpCmd = &cobra.Command{
	Use:   "statsDump",
	Short: "Export statistics and table structures",
	Long:  `Export statistics and table structures`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("statsDump called")
		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", dbusername, dbpassword, dbhost, dbport, dbname)
		command.Show(dsn)
	},
}

func init() {
	rootCmd.AddCommand(statsDumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statsDumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statsDumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	statsDumpCmd.Flags().StringVarP(&dbusername, "dbusername", "u", "root", "Database user")
	statsDumpCmd.Flags().StringVarP(&dbname, "dbname", "d", "test", "Database name")
	statsDumpCmd.Flags().StringVarP(&dbhost, "dbhost", "H", "127.0.0.1", "Database host")
	statsDumpCmd.Flags().StringVarP(&dbpassword, "dbpassword", "p", "", "Database passowrd")
	statsDumpCmd.Flags().IntVarP(&dbport, "dbport", "P", 4000, "Database port")
}
