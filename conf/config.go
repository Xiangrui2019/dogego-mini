package conf

import (
	"dogego-mini/cache"
	"dogego-mini/executers"
	"dogego-mini/models"
	"dogego-mini/modules"
	"dogego-mini/tasks"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	models.ConnectDatabase(os.Getenv("DATABASE_DSN"))
	cache.ConnectRedisCache()
	modules.InitAllModules()
	tasks.StartCronJobs(false)
	executers.TimeExecuter()
	executers.AsyncExecuter()
}
