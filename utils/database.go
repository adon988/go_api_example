package utils

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
	viper "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client
var o sync.Once
var db *gorm.DB

type InfoDbHost struct {
	Host     string
	Port     any
	Username string
	Password string
	Database string
}

func (infoDb InfoDbHost) InitDB() (*gorm.DB, error) {
	var err error
	o.Do(func() {
		err = mapstructure.Decode(viper.GetStringMap("mysql"), &infoDb)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", infoDb.Username, infoDb.Password, infoDb.Host, infoDb.Port, infoDb.Database)
		fmt.Println("Init DB once")
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
