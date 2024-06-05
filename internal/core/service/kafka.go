package service

import (
	"encoding/json"
	"pos/internal/core/domain"
	"pos/internal/core/port"
)

// type KafkaService interface {
// 	// ReadMessage(newsRepo repo.NewsRepository, elasticRepo repo.ElasticRepository) error // TEMP
// }

type KafkaService struct {
	repo port.KafkaRepository
}

func NewKafkaService(repo port.KafkaRepository) *KafkaService {
	return &KafkaService{
		repo,
	}
}

// TEMP
// func (ks *KafkaService) ReadMessage(newsRepo repo.NewsRepository, elasticRepo repo.ElasticRepository) error {
func (ks *KafkaService) ReadMessage() error {
	dataChan := make(chan []byte) // it will be sent to ReadMessage function

	go func() {
		for {
			select {
			case dataByte := <-dataChan:
				data := new(domain.Event)
				if e := json.Unmarshal(dataByte, data); e != nil {
					return
				}
				// elasticData := m.ElasticNews{
				// 	ID:      data.ID,
				// 	Created: data.Created,
				// }
				// if e := elasticRepo.Store(elasticData); e != nil {
				// 	return
				// }

				// if e := newsRepo.Store(data); e != nil {
				// 	return
				// }
			default:
			}
		}
	}()

	ks.repo.ReadMessage(dataChan)

	return nil
}
