package etcddemo

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
)

func Register() {
	rsp,err := etcdCli.Put(context.TODO(),"/test/b1","something",clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
}
