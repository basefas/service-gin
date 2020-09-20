package main

import (
	"basefas.com/service-gin/cmd/app/handlers"
	"basefas.com/service-gin/internal/auth"
	"basefas.com/service-gin/internal/utils/conf"
	"basefas.com/service-gin/internal/utils/db/mysql"
)

func main() {
	run()
}

func run() {
	conf.Init()
	mysql.Init()
	auth.Init()
	handlers.Init()
}
