package nsqsingle

import (
	"goTool/nsqdemo"
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

func Product() {
	conf := nsq.NewConfig()
	product, err := nsq.NewProducer(nsqdemo.Addr, conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 5; i++ {
		message := "product message :" + time.Now().Format("2006-01-02 15:04:05")
		//fmt.Println(message)
		// 发送消息
		product.Publish(nsqdemo.Topic, []byte(message))
		time.Sleep(time.Second)
	}

}
