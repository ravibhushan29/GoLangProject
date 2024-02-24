package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
	db *gorm.DB
)

func Connnect() {
	dsn := "host=localhost user=your_username password=your_password dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

}

