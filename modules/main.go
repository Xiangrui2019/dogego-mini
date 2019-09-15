package modules

func InitAllModules() {
	InitLockerModule()
	InitHealthChecksModule()
	InitSentryClient()
}
