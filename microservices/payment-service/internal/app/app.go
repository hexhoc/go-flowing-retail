package app

import (
	"context"

	"github.com/hexhoc/payment-service/config"
	"github.com/hexhoc/payment-service/pkg/kafka/consumer"
	"github.com/hexhoc/payment-service/pkg/kafka/publisher"
	"github.com/hexhoc/payment-service/pkg/logger"
	log "github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) {
	// LOGGER INIT
	logger.Init(cfg.LogLevel)
	log.Info("Starting order-service")

	// KAFKA PUBLISHER INIT
	paymentEventPublisher := publisher.NewPublisher([]string{cfg.KafkaAddr}, "paymentReceivedTopic")
	// KAFKA CONSUMER
	consumer.Consume(context.Background(), []string{cfg.KafkaAddr}, "paymentRetrieveTopic", messageListener.EventPayment)

}
