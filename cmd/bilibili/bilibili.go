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
		docs.SwaggerInfo.Description = "Bili-Bili API Example by Caesar"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "fp-sanbercode-go-45-caesar-production.up.railway.app"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := router.SetupRouter(db)
	r.Run()
}
