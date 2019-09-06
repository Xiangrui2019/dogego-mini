package conf

import (
	"dogego-mini/cache"
	"dogego-mini/models"
	"dogego-mini/modules"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	models.ConnectDatabase(os.Getenv("DATABASE_DSN"))
	cache.ConnectRedisCache()
	modules.InitAllModules()
}
