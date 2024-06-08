package utils

import (
	"github.com/adon988/go_api_example/docs"
)

func InitSwagger() {
	docs.SwaggerInfo.Title = Configs.Doc.Title
	docs.SwaggerInfo.Description = Configs.Doc.Description
	docs.SwaggerInfo.Version = Configs.Doc.Version
	docs.SwaggerInfo.Host = Configs.Doc.Host
	docs.SwaggerInfo.BasePath = Configs.Doc.BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
