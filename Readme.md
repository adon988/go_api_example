## Getting start

In this gorm example project, you should setting the default configuration and environment by following:

### Config database connection

At `config/config.yml.example` is a configuration example, need to rename this example to `config/config.yml` to setting database connection.

```
cp confing/config.yml.example config/config.yml
```


After add config.yml, If you wanna change setting config from environment variables, you can do that:

```
export MYSQL_USERNAME=root
export MYSQL_PASSWORD=root
export MYSQL_DATABASE=local_test
export MYSQL_PORT=3306
export MYSQL_HOST=localhost
export JWT_SECRET="jwt secret string"
export JWT_EXPIRE_AT=7200
export JWT_EFFECT_AT=-1000
export JWT_ISSUER=adam
```

The environment variables will auto replace config value.

### Initialize go env

Directly install the go versions by [asdf](https://hoohoo.top/blog/20240315145428-asdf-quick-note/) .tool-versions with:

```
asdf install
```

### Initialize Table

Run the project
```
go mod tidy
go run main.go
```


### DB Migration

The basic models (internal/models/basic.go) are migration mappings and can be migrated by following the migration process.

Migrate all tables

```
go run cmd/migrate/main.go -automigrate
```

Migrate single table

```
go run cmd/migrate/main.go -migrate_table=Organization
```


### Try APIs

Get JWT token
```
curl --location 'localhost:8080/jwt_token'
```
Fetch member data
```
curl --location 'localhost:8080/member/11' \
--header 'Authorization: Bearer {JWT_TOKEN}
```

### Swagger

Reference https://github.com/swaggo/swag

Local install the swag
```
go get -u github.com/swaggo/swag/cmd/swag

//or Go 1.17
go install github.com/swaggo/swag/cmd/swag@latest

```
Install swag
```
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
//or fish
set -U fish_user_paths (go env GOPATH)/bin $fish_user_paths
```

go get

```
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger
```

generate doc (swagger.json, swagger.yaml) to `go-project-name/docs.`

```
cd cmd/server/

swag init --dir ./cmd/server,./internal/ --output ./docs


//or execute by makefile
make swag_init
```

add swagger access page

```
swaggerfiles "github.com/swaggo/files"
ginSwagger "github.com/swaggo/gin-swagger"

r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
```

Definication the swagger document path

```

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
```

### DB cluster (read/write) resolver

https://gorm.io/docs/dbresolver.html

utils/database.go 可使用 database connection resolver

### Casbin

The casbin permission control only used on small user and role case (like backsite, cms), the benchmarks only suitable for users < 10k.

Install casbin for permission control
```
go get github.com/casbin/casbin/v2
```

Gorm adapter to save the policy to DB

```
go get github.com/casbin/gorm-adapter/v3
```

