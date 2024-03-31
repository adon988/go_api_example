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

func InitDB() (*gorm.DB, error) {

	host := viper.GetString("mysql.localhost")
	port := viper.GetString("mysql.port")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	o.Do(func() {
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
