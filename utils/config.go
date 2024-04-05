package utils

import (
	"fmt"
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

	// viper auto read all env variables, the key will auto uppercase
	viper.AutomaticEnv()
	//Set prefix of env variables, for example "MYE_" will be used `viper.SetEnvPrefix("MYE")`
	viper.SetEnvPrefix("")

	//Replace the environment variables _ to .
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
