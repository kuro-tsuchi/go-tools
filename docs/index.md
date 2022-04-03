# 1. tools

## 1.1. hotDeployment

```bash
go install github.com/cosmtrek/air@latest
air
```

## 1.2. devTools

### 1.2.1. sql2go

用于将 sql 语句转换为 golang 的 struct. 使用 ddl 语句即可。
`show create table xxx.`

### 1.2.2. toml2go

用于将编码后的 toml 文本转换问 golang 的 struct.
<https://xuri.me/toml-to-go/>

### 1.2.3. curl2go

用来将 curl 命令转化为具体的 golang 代码.
<https://mholt.github.io/curl>-...

### 1.2.4. json2go

用于将 json 文本转换为 struct.
<https://mholt.github.io/json>-...

### 1.2.5. mysql 转 ES 工具

<http://www.ischoolbar.com/EsP>...

### 1.2.6. golang

模拟模板的工具，在支持泛型之前，可以考虑使用。
<https://github.com/cheekybits>...

### 1.2.7. 查看某一个库的依赖情况，类似于 go list 功能

<https://github.com/KyleBanks/depth>

### 1.2.8. 一个好用的文件压缩和解压工具，集成了 zip，tar 等多种功能，主要还有跨平台

<https://github.com/mholt/archiver>

### 1.2.9. go 内置命令

go list 可以查看某一个包的依赖关系.
go vet 可以检查代码不符合 golang 规范的地方。

### 1.2.10. 热编译工具

<https://github.com/cosmtrek/air>
<https://github.com/silenceper/gowatch>

### 1.2.11. revive

golang 代码质量检测工具
<https://github.com/mgechev/revive>

### 1.2.12. Go Callvis

golang 的代码调用链图工具
<https://github.com/ofabry/go-callvis>

### 1.2.13. Realize

开发流程改进工具
<https://github.com/oxequa/realize>

### 1.2.14. Gotests

自动生成测试用例工具
<https://github.com/cweill/gotests>

## 1.3. 调试工具

### 1.3.1. perf

代理工具，支持内存，cpu，堆栈查看，并支持火焰图.
perf 工具和 go-torch 工具，快捷定位程序问题.
<https://github.com/uber-archi>...
<https://github.com/google/gops>

### 1.3.2. dlv 远程调试

基于 goland+dlv 可以实现远程调式的能力.
<https://github.com/go-delve/d>...
提供了对 golang 原生的支持，相比 gdb 调试，简单太多。

### 1.3.3. 网络代理工具

goproxy 代理，支持多种协议，支持 ssh 穿透和 kcp 协议.
<https://github.com/snail007/g>...

### 1.3.4. 抓包工具

go-sniffer 工具，可扩展的抓包工具，可以开发自定义协议的工具包. 现在只支持了 http，mysql，redis，mongodb.
基于这个工具，我们开发了 qapp 协议的抓包。
<https://github.com/40t/go-sniffer>

### 1.3.5. 反向代理工具，快捷开放内网端口供外部使用

ngrok 可以让内网服务外部调用
<https://ngrok.com/>
<https://github.com/inconshrev>...

### 1.3.6. 配置化生成证书

从根证书，到业务侧证书一键生成.
<https://github.com/cloudflare>...

### 1.3.7. 免费的证书获取工具

基于 acme 协议，从 letsencrypt 生成免费的证书，有效期 1 年，可自动续期。
<https://github.com/Neilpang/a>...

### 1.3.8. 开发环境管理工具，单机搭建可移植工具的利器。支持多种虚拟机后端

vagrant常被拿来同 docker 相比，值得拥有。
<https://github.com/hashicorp/>...

### 1.3.9. 轻量级容器调度工具

nomad 可以非常方便的管理容器和传统应用，相比 k8s 来说，简单不要太多.
<https://github.com/hashicorp/>...

### 1.3.10. 敏感信息和密钥管理工具

<https://github.com/hashicorp/>...

### 1.3.11. 高度可配置化的 http 转发工具，基于 etcd 配置

<https://github.com/gojek/weaver>

### 1.3.12. 进程监控工具 supervisor

<https://www.jianshu.com/p/39b>...

### 1.3.13. 基于procFile进程管理工具. 相比 supervisor 更加简单

<https://github.com/ddollar/fo>...

### 1.3.14. 基于 http，https，websocket 的调试代理工具，配置功能丰富。在线教育的 nohost web 调试工具，基于此开发

<https://github.com/avwo/whistle>

### 1.3.15. 分布式调度工具

<https://github.com/shunfei/cr..._ZH.md>
<https://github.com/ouqiang/go>...

### 1.3.16. 自动化运维平台 Gaia

<https://github.com/gaia-pipel>...

## 1.4. 网络工具

![20220401212704](<https://raw.githubusercontent.com/kuro-tsuchi/my-picgo/master/md/img/20220401212704.png?token=ANWZ3QOTKWNR5NGJVQOND5TCI366I>###

## 1.5. 常用网站

 1. go 百科全书: <https://awesome-go.com/>
 1. json 解析: <https://www.json.cn/>
 1. 出口 IP: <https://ipinfo.io/>
 1. redis 命令: <http://doc.redisfans.com/>
 1. ES 命令首页  <https://www.elastic.co/guide/>
 1. UrlEncode: <http://tool.chinaz.com/Tools/>
 1. Base64: <https://tool.oschina.net/encr>
 1. Guid: <https://www.guidgen.com/>
 1. 常用工具: <http://www.ofmonkey.com/>

## 1.6. golang 常用库

### 1.6.1. 日志

<https://github.com/Sirupsen/logrus>
<https://github.com/uber-go/zap>

#### 1.6.1.1. 配置

兼容 json，toml，yaml，hcl 等格式的日志库.
<https://github.com/spf13/viper>

### 1.6.2. 存储

1. mysql: <https://github.com/go-xorm/xorm>
1. es: <https://github.com/elastic/elasticsearch>
1. redis: <https://github.com/go-redis/redis>  <https://github.com/gomodule/redigo>
1. mongo: <https://github.com/mongodb/mongo-go-driver>
1. kafka:  <https://github.com/Shopify/sarama>

### 1.6.3. 数据结构

<https://github.com/emirpasic/gods>

### 1.6.4. 命令行

<https://github.com/spf13/cobra>

### 1.6.5. 框架

<https://github.com/grpc/grpc-go>
<https://github.com/gin-gonic/gin>

### 1.6.6. 并发

<https://github.com/Jeffail/tunny>
<https://github.com/benmanns/goworker>
现在我们框架在用的，虽然 star 不多，但是确实好用，当然还可以更好用.
<https://github.com/rafaeldias>...

### 1.6.7. 工具

定义了实用的判定类，以及针对结构体的校验逻辑，避免业务侧写复杂的代码.
<https://github.com/asaskevich>...
<https://github.com/bytedance/>...
protobuf 文件动态解析的接口，可以实现反射相关的能力。
<https://github.com/jhump/prot>...

### 1.6.8. 表达式引擎工具

<https://github.com/Knetic/gov>...
<https://github.com/google/cel-go>

### 1.6.9. 字符串处理

<https://github.com/huandu/xstrings>

### 1.6.10. ratelimit 工具

<https://github.com/uber-go/ra>...
<https://blog.csdn.net/chencho>...
<https://github.com/juju/ratel>...

### 1.6.11. golang 熔断的库

熔断除了考虑频率限制，还要考虑 qps，出错率等其他东西.

<https://github.com/afex/hystrix-go>
<https://github.com/sony/gobreaker>

### 1.6.12. 表格

<https://github.com/chenjiando>...
<https://github.com/go-echarts/go-echarts>

### 1.6.13. tail 工具库

<https://github.com/hpcloud/tail>
