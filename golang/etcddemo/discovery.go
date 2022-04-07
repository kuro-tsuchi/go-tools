package etcddemo

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Discovery() {
	GetInfo()
}
func GetInfo() {
	//设置OpOption，前缀查询
	rsp, err := etcdCli.Get(context.TODO(), "/test/b", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
	}

	for _, kv := range rsp.Kvs {
		fmt.Println("key=", string(kv.Key), "|", "value=", string(kv.Value))
	}
}