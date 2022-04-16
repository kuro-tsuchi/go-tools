package gormdemo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MainExec() {
	dsn := "root:root123@tcp(127.0.0.1:3306)/golang2022?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d, _ := db.DB()
	fmt.Printf("d.Ping(): %v\n", d.Ping())

}
