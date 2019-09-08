package healthchecks

import "dogego-mini/modules"

func RegisterHealthChecks() {
	modules.HealthChecksModule.AddHealthCheck(DatabaseHealthCheck)
	modules.HealthChecksModule.AddHealthCheck(RedisHealthCheck)
}
