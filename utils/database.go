package utils

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Configs.Mysql.Username, Configs.Mysql.Password, Configs.Mysql.Host, Configs.Mysql.Port, Configs.Mysql.Database)
		fmt.Println("Init DB once", dsn)
		db, err = func() (*gorm.DB, error) {
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("Failed to connect to database: " + err.Error())
			}
			fmt.Println("Connection to database!")
			return db, err
		}()
	})
	return db, err
}
