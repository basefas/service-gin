package main

import (
	"basefas.com/service-gin/cmd/app/handlers/v1"
	"basefas.com/service-gin/internal/utils/conf"
	"basefas.com/service-gin/internal/utils/db/mysql"
)

func main() {
	run()
}

func run() {
	conf.Init()
	mysql.Init()
	v1.Init()
}
