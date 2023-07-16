package bilibili

import (
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/v1/router"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/docs"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"
)

func init() {
	environment := util.GetEnv("NODE_ENV", "development")

	if environment == "development" {
		//programmatically set swagger info
		docs.SwaggerInfo.Title = "Bili-Bili Marketplace API"
		docs.SwaggerInfo.Description = "Admin Register x-api-key = 9c6f9769-6d5b-493d-ae2e-4fad70711564"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := router.SetupRouter(db)
	r.Run()
}
