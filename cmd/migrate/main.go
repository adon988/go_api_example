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

func init() {
	flag.BoolVar(&runAutoMigrations, "automigrate", false, "run auto migrations")
	flag.StringVar(&runMigrateTable, "migrate_table", "", "run migration table")
	flag.Parse()

}
func main() {

	if !runAutoMigrations && runMigrateTable == "" {
		fmt.Println("Please use -automigrate or -migrate_table=TableName")
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
}
