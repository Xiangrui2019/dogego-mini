package healthchecks

import "dogego-mini/models"

func DatabaseHealthCheck() error {
	return models.DB.DB().Ping()
}
