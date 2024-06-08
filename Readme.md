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


### Initialize Table

Run the project and auto migration to generate `members` table and execute CRUD 

```
go mod tidy
go run main.go
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
swag init
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