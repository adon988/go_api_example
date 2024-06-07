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