package main

import (
	"fmt"
	"time"
	"unsafe"
)

type ProducerMessage struct {
	// Topic
	Topic string

	// 数据（发送给消费者）
	Value []byte

	// k-v 数据 （发送给消费者）
	Headers []RecordHeader

	// 自定义，不会发送给消费者.
	Metadata unsafe.Pointer

	// 消息序号
	Offset int64

	// 消息在传输系统中成功写入的时间戳
	ConfirmTimestamp time.Time
}

// RecordHeader ...
type RecordHeader struct {
	Key   []byte `json:"Key"`
	Value []byte `json:"Value"`
}

type ProducerError struct {
	Msg *ProducerMessage
	Err error
}

type ProduceHandler interface {
	OnProduceSuccess(producedMsg *ProducerMessage)
	OnProduceError(producerErr *ProducerError)
}

type MsgHandler struct {
}

func (s *MsgHandler) OnProduceSuccess(producedMsg *ProducerMessage) {
	fmt.Printf("sent msg, topic:%v, offset:%v\n", producedMsg.Topic, producedMsg.Offset)
}

func (s *MsgHandler) OnProduceError(producerErr *ProducerError) {
	panic(producerErr.Err)
}

func main(){
	mh := &MsgHandler{}

	msg := &ProducerMessage{
		Topic: "foo",
		Value: []byte("bar"),
		Metadata: unsafe.Pointer(mh),
	}

	if msg.Metadata != nil {
		fmt.Printf("meta:%p\n", msg.Metadata)
		fmt.Printf("xxx:%p\n", mh)

		iii := (*ProduceHandler)(msg.Metadata)
		fmt.Printf("iii:%p\n", iii)

		v := *iii
		fmt.Printf("vvv:%p\n", &v)

		// v.OnProduceSuccess(msg) // panic

		aaa := (*interface{})(msg.Metadata)
		fmt.Printf("aaa:%p\n", aaa)

		bbb := *aaa
		fmt.Printf("bbb:%p\n", bbb)

		ccc := bbb.(ProduceHandler) // panic
		fmt.Printf("ccc:%p\n", ccc)
	}
}
