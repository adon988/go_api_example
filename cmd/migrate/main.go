package main

import (
	"flag"
	"fmt"

	"github.com/adon988/go_api_example/internal/migrations"
	"github.com/adon988/go_api_example/internal/utils"
)

var runAutoMigrations bool
var runMigrateTable string
var InfoDb utils.InfoDb
var runAlertTable string

func init() {
	flag.BoolVar(&runAutoMigrations, "automigrate", false, "run auto migrations")
	flag.StringVar(&runMigrateTable, "migrate_table", "", "run migrate table")
	flag.StringVar(&runAlertTable, "alert_table", "", "run alert table")
	flag.Parse()

}
func main() {

	if !runAutoMigrations && runMigrateTable == "" && runAlertTable == "" {
		fmt.Println("Please use -automigrate or -migrate_table=TableName or -alert_table=AlertTableName")
		return
	}

	// Init config
	err := utils.InitConfig()
	if err != nil {
		panic("init config error:" + err.Error())
	}

	Db, _ := InfoDb.InitDB()
	Migration := migrations.NewMigration{Db: Db}
	// Migrate all tables (ex. go run cmd/migrate/main.go -automigrate)
	if runAutoMigrations {
		Migration.AutoMigration()
		Migration.AutoMigrateCasbin()
		fmt.Println("Auto migrate all tables success")
	}

	// Migrate single table (ex. go run cmd/migrate/main.go -migrate_table=Organization)
	if res := Migration.MigrationTable(runMigrateTable); res {
		fmt.Println("Migrate table " + runMigrateTable + " success")
	}

	if res := Migration.AlertTable(runAlertTable); res {
		fmt.Println("Alert table " + runAlertTable + " success")
	}
}
