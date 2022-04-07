package main

import (
	"github.com/sirupsen/logrus"
	"goTool/utils/redislock"
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
	redislock.MainExec()
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
