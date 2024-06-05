package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"pos/internal/core/domain"
	"pos/internal/core/port"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type KafkaAdapter struct {
	conn    *kafka.Conn
	url     string
	topic   string
	timeout time.Duration
}

func KafkaConnection(timeout, url, topic string) port.KafkaRepository {
	timeoutInt, _ := strconv.Atoi(timeout)
	if timeoutInt == 0 {
		timeoutInt = 10
	}

	if url == "" {
		url = "localhost:9092"
	}

	if topic == "" {
		topic = "event"
	}

	repo, e := NewKafkaConnection(url, topic, timeoutInt)
	if e != nil {
		log.Fatal(e) // TEMP
	}

	return repo
}

func newKafkaConnection(URL, topic string) (*kafka.Conn, error) {

	// TEMP
	fmt.Println("KAFKA URL: ", URL)
	fmt.Println("KAFKA TOPIC: ", topic)

	kafkaConn, e := kafka.DialLeader(context.Background(), "tcp", URL, topic, 0)
	if e != nil {
		return nil, errors.Wrap(e, "repository.newKafkaConnection")
	}

	return kafkaConn, e
}

func NewKafkaConnection(URL, topic string, timeout int) (port.KafkaRepository, error) {
	ka := &KafkaAdapter{
		topic:   topic,
		url:     URL,
		timeout: time.Duration(timeout) * time.Second,
	}

	conn, e := newKafkaConnection(URL, topic)
	if e != nil {
		return nil, errors.Wrap(e, "repository.NewKafkaConnection")
	}

	ka.conn = conn

	return ka, nil
}

func (k KafkaAdapter) WriteMessage(data *domain.Event) error {
	msgs, e := json.Marshal(data)
	if e != nil {
		return e
	}

	k.conn.SetWriteDeadline(time.Now().Add(k.timeout))

	if _, e = k.conn.WriteMessages(
		kafka.Message{Value: msgs},
	); e != nil {
		return e
	}
	return nil
}

func (k KafkaAdapter) ReadMessage(res chan<- []byte) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{k.url},
		Topic:     k.topic,
		Partition: 0,
		MinBytes:  10,
		MaxBytes:  10e3,
	})
	ctx := context.Background()
	lastOffset, _ := k.conn.ReadLastOffset() // get latest offset
	r.SetOffset(lastOffset)                  // set latest offset

	for {
		m, e := r.ReadMessage(ctx)
		if e != nil {
			log.Println("kafka-repo ReadMessage", e.Error())
			break
		}
		// fmt.Printf("message at offset %d: %s = %s at %v\n", m.Offset, string(m.Key), string(m.Value), m.Time)
		res <- m.Value
	}

}
