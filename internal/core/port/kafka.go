package port

import "pos/internal/core/domain"

type KafkaService interface {
	// WriteMessage(data *domain.Event) error
	ReadMessage(res chan<- []byte)
}

type KafkaRepository interface {
	WriteMessage(data *domain.Event) error
	ReadMessage(res chan<- []byte)
}
