package utils

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB
var RedisClient *redis.Client
var o sync.Once
var db *gorm.DB

type InfoDb struct {
	Mysql Mysql
}

func (infoDb InfoDb) InitDB() (*gorm.DB, error) {
	var err error
	o.Do(func() {
		dsnPrimary := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Configs.Db.Writer.Username, Configs.Db.Writer.Password, Configs.Db.Writer.Host, Configs.Db.Writer.Port, Configs.Db.Writer.Database)
		dsnReader := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Configs.Db.Reader.Username, Configs.Db.Reader.Password, Configs.Db.Reader.Host, Configs.Db.Reader.Port, Configs.Db.Reader.Database)
		fmt.Println("Init DB once", dsnPrimary)
		db, err = func() (*gorm.DB, error) {
			db, err := gorm.Open(mysql.Open(dsnPrimary), &gorm.Config{}) // db1 Reader
			if err != nil {
				panic("Failed to connect to database: " + err.Error())
			}
			db.Use(dbresolver.Register(dbresolver.Config{
				Sources: []gorm.Dialector{mysql.Open(dsnReader)}, // db2 reader
				// Replicas: []gorm.Dialector{mysql.Open(dsn)}, // optional db3, db4

			}))
			// Debug mode
			if Configs.Db.Debug_Mode {
				db = db.Debug()
			}
			fmt.Println("Connection to database!")
			return db, err
		}()

	})
	return db, err
}
