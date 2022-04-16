package etcddemo

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var  etcdCli *clientv3.Client

func init() {
	var err error
	etcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	//defer etcdCli.Close()
	if err != nil {
		log.Fatal(err)
	}
}
