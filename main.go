package main

import (
	"goTool/gormdemo"

	"github.com/sirupsen/logrus"
)

func main() {
	// sql2go
	//sql2gorm.MainExec(test.SqlMock)

	//etcd
	//etcddemo.Register()
	//etcddemo.Discovery()

	//nsq 单机
	//nsqsingle.Product()
	//nsqsingle.Consumer()

	////nsq 集群
	//nsqcluster.Product()
	//nsqcluster.Consumer()

	// 雪花算法
	//snowflake.MainExec()
	// redislock.MainExec()

	// gorm mysql

	gormdemo.MainExec()
	// gorm pgsql
	//pgsqldemo.MainExec()

	// ckdemo.MainExec()
	//mysqlDemo.MainExec()
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
