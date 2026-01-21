package kafka

import (
	"context"
	"os"
	"time"
	"xplr-distributed-mq/mq"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Config struct {
	Brokers      []string
	KafkaVersion string
	MaxRetry     int

	AllowAutoTopicCreation bool
	EnableDebug            bool
}

type kafkaGoProducer struct {
	conf     *Config
	producer *kgo.Client
}

func NewProducer(conf *Config) (mq.Producer, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(conf.Brokers...),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.RecordRetries(conf.MaxRetry),
		kgo.RecordDeliveryTimeout(10 * time.Second),
	}

	if conf.AllowAutoTopicCreation {
		opts = append(opts, kgo.AllowAutoTopicCreation())
	}

	logLevel := kgo.LogLevelWarn
	if conf.EnableDebug {
		logLevel = kgo.LogLevelDebug
	}

	opts = append(opts,
		kgo.WithLogger(
			kgo.BasicLogger(os.Stdout, logLevel, func() string {
				return "kafka-producer"
			}),
		),
	)

	cl, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return &kafkaGoProducer{
		conf:     conf,
		producer: cl,
	}, nil
}

func (p *kafkaGoProducer) Publish(ctx context.Context, msg *mq.Message) error {
	record := &kgo.Record{
		Topic: msg.Topic,
		Key:   []byte(msg.Key),
		Value: msg.Value,
	}

	if err := p.producer.ProduceSync(ctx, record).FirstErr(); err != nil {
		return err
	}

	return nil
}

func (p *kafkaGoProducer) Close() error {
	p.producer.Close()

	return nil
}
