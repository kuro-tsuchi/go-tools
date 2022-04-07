package nsqsingle

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"goTool/example/nsqdemo"
	"log"
)

type NewHandler struct{}

func (m *NewHandler) HandleMessage(msg *nsq.Message) (err error) {
	addr := msg.NSQDAddress
	message := string(msg.Body)
	fmt.Println(addr, message)
	return
}
func MyConsumers(topic, channel, addr string) {
	conf := nsq.NewConfig()
	fmt.Printf("channel=%v\n", channel)
	newConsumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		log.Fatalln(err)
	}

	// 接收消息
	//newHandler := &NewHandler{}
	//newConsumer.AddHandler(newHandler)

	newConsumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fmt.Println( channel +" " +  string(msg.Body))
		return nil
	}))
	err = newConsumer.ConnectToNSQD(addr)
	if err != nil {
		log.Fatalln(err)
	}
}
func Consumer() {
	// 模拟多个从多个channel去消息
	go MyConsumers(nsqdemo.Topic, "channel-aa", nsqdemo.Addr)
	go MyConsumers(nsqdemo.Topic, "channel-bb", nsqdemo.Addr)
	select {}
}
