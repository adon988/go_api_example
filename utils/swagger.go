package utils

import (
	"github.com/adon988/go_api_example/docs"
)

func InitSwagger() {
	docs.SwaggerInfo.Title = "API Document"
	docs.SwaggerInfo.Description = "Swagger API document"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
