# nsq

## 单机

1. nsqd 启动 nsq服务, tcp: 4150, http:4151
2. 启动 web界面,  nsqadmin --nsqd-http-address="127.0.0.1:4151"
3. web界面:  http://127.0.0.1:4171/

## 集群
1. 启动nsqlookupd , nsqlookupd是nsq管理集群拓扑信息以及用于注册和发现nsqd服务
2. 添加nsqd 实例, nsqd -http-address="0.0.0.0:8080" -tcp-address="0.0.0.0:8081" -lookupd-tcp-address="127.0.0.1:4160", nsqd -http-address="0.0.0.0:8090" -tcp-address="0.0.0.0:8091" -lookupd-tcp-address="127.0.0.1:4160"

## 启动nsqadmin
nsqadmin -lookupd-http-address="127.0.0.1:4161"
##  浏览器中访问nsqadmin

http://127.0.0.1:4171