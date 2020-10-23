// Copyright © 2020 NAME HERE <EMAIL ADDRESS>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/WalterWj/go-study/imp"

	_ "github.com/go-sql-driver/mysql" // import mysql
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	name    string
	age     int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-study",
	Short: "A test demo",
	Long:  `Demo is a test appcation for print things.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			_ = cmd.Help()
			return
		}
		imp.Show(name, age)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name ")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-study" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-study")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
func mysqlConnect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	ifErrWithLog(err)
	return db
}

func ifErrWithLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
