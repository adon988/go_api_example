package utils

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	viper "github.com/spf13/viper"
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

		if err := viper.Unmarshal(&infoDb); err != nil {
			panic(fmt.Errorf("unable to decode into struct, %v", err))
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", infoDb.Mysql.Username, infoDb.Mysql.Password, infoDb.Mysql.Host, infoDb.Mysql.Port, infoDb.Mysql.Database)
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
