package api

import (
	log "github.com/sirupsen/logrus"
	mydb "goTool/gormdemo/db"
	"goTool/gormdemo/modules"
)

// 获取第一条记录（主键升序）
func First() (*modules.Product, error) {
	dest := &modules.Product{}
	err := db.First(dest).Error
	if err != nil {
		log.Error(err.Error())
		return dest, nil
	}
	return dest, err
}

// Last 获取最后一条记录（主键降序）
func Last() *modules.Product {
	dest := &modules.Product{}
	db.Last(dest)
	return dest
}

// 获取一条记录，没有指定排序字段
func Take() *modules.Product {
	dest := &modules.Product{}
	db.Take(dest)
	return dest

}

func init() {
	db = mydb.InitDB()
	product := &modules.Product{}
	db.AutoMigrate(product)
}
