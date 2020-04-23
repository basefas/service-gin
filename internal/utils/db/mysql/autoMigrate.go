package mysql

import (
	"basefas.com/service-gin/internal/user"
	"basefas.com/service-gin/internal/utils/db"
)

func AutoMigrate() {
	db.Mysql.AutoMigrate(&user.User{})
}
