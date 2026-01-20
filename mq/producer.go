package mq

import "context"

type Message struct {
	Topic string
	Key   string
	Value []byte
}

type Producer interface {
	Publish(ctx context.Context, msg *Message) (partition int32, offset int64, err error)

	Close() error
}
