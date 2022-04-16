package nsqcluster


import (
    "bufio"
    "fmt"
    "github.com/nsqio/go-nsq"
    "log"
    "os"
    "strings"
)

var pro *nsq.Producer

func NewPro(addr string) (err error) {
    conf := nsq.NewConfig()
    pro, err = nsq.NewProducer(addr, conf)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}
func Product() {
    nsqAddr := "127.0.0.1:8091"
    err := NewPro(nsqAddr)
    if err != nil {
        fmt.Println(err)
        return
    }else{
        fmt.Println("connect 127.0.0.1:8091 success")
    }
    // 读取标准输入
    reader := bufio.NewReader(os.Stdin)
    for {
        // 读取所有内容直到遇见回车(\n)
        data, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("read data from stdin is field : ", err)
            continue
        }
        // 当输入q的时候退出
        data = strings.TrimSpace(data)
        if strings.ToUpper(data) == "Q" {
            break
        }
        err = pro.Publish("topic-demo1", []byte(data))
        if err != nil {
            fmt.Println("nsq publish is field ", err)
            continue
        }
    }
    fmt.Println("exit !")
}