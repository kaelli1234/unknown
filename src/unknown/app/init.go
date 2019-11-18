package app

import (
	"ccgwf/app"
)

func Init(cfgPath *string) {
	app.Init(cfgPath)
	Kernel().SetConfig(*cfgPath)
}

func initApp() {
}
