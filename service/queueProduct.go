package service

import (
	"app/config"
	"app/model"
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type queueProductService struct {
	channelQueueProduct *amqp091.Channel
}

type QueueProductService interface {
	PushMessInQueueToElasticSearch(data map[string]interface{}) error
}

func (s *queueProductService) PushMessInQueueToElasticSearch(data map[string]interface{}) error {
	dataBytes, errConvert := json.Marshal(data)
	if errConvert != nil {
		return errConvert
	}

	errPush := s.channelQueueProduct.PublishWithContext(context.Background(),
		"",
		string(model.PRODUCT_TO_ELASTIC),
		false, // mandatory
		false, // immediate,
		amqp091.Publishing{
			ContentType:  "text/plain",
			Body:         dataBytes,
			DeliveryMode: amqp091.Persistent,
		},
	)

	if errPush != nil {
		return errPush
	}
	return nil
}

func NewQueueProductService() QueueProductService {
	return &queueProductService{
		channelQueueProduct: config.GetRabbitChannel(),
	}
}
