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

type Db struct {
	Writer     Mysql
	Reader     Mysql
	Debug_Mode bool
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
type Jwt struct {
	Secret    string
	Expire_At int64
	Effect_At int64
	Issuer    string
}
type Doc struct {
	Title       string
	Description string
	Version     string
	Host        string
	BasePath    string
}
type Gin struct {
	Debug_Mode    bool
	Port          int
	Read_Timeout  int
	Write_Timeout int
}
type ACM_MODEL struct {
	Config string
}
type ACM struct {
	Model ACM_MODEL
}

type Service struct {
	Page_Items int
}
type Config struct {
	App     App
	Db      Db
	Redis   Redis
	Jwt     Jwt
	Doc     Doc
	Gin     Gin
	ACM     ACM
	Service Service
}

var Configs Config

func InitConfig() error {

	//set config file as default
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config/")

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

	if err := viper.Unmarshal(&Configs); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	// fmt.Printf("MYSQL_USERNAME: %s\n", Configs.Mysql.Username)
	// fmt.Printf("MYSQL_PASSWORD: %s\n", Configs.Mysql.Password)
	// fmt.Printf("JWT: %s\n", Configs.Jwt.Secret)
	// fmt.Printf("JWT: %d\n", Configs.Jwt.Expire_At)
	return err
}
