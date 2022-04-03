package main

import (
	"github.com/sirupsen/logrus"
	"goTool/utils/xstrings"
)

func main() {
	// sql2go
	//sql2gorm.MainExec(sql)
	//toml2go

	xstrings.MainExec()

}

var sql = `CREATE TABLE person_info (
	id BIGINT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT 'primary id',
	name VARCHAR(30) NOT NULL DEFAULT 'default_name',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	age INT(11) unsigned NULL,
	sex VARCHAR(2) NULL
	);`

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
