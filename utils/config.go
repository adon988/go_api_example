package utils

import (
	"fmt"
	"os"
	"strings"

	viper "github.com/spf13/viper"
)

type Mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}
type App struct {
	Name   string
	Prefix string
}
type Redis struct {
	Host        string
	Port        string
	Password    string
	SelectDb    int
	PolSize     int
	MinIdleConn int
}
type Config struct {
	App   App
	Mysql Mysql
	Redis Redis
}

func InitConfig() error {

	//set config file as default
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	// Set environment variables
	os.Setenv("MYSQL_USERNAME", "root")
	os.Setenv("MYSQL_PASSWORD", "root")
	os.Setenv("MYSQL_DATABASE", "local_test")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_HOST", "localhost")

	// viper auto read all env variables, the key will auto uppercase
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// var config Config
	// if err := viper.Unmarshal(&config); err != nil {
	// 	panic(fmt.Errorf("unable to decode into struct, %v", err))
	// }
	// fmt.Printf("MYSQL_USERNAME: %s\n", config.Mysql.Username)
	// fmt.Printf("MYSQL_PASSWORD: %s\n", config.Mysql.Password)
	return err
}
