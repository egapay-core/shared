package broker

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/egapay-core/shared/vault"
	"log"
)

type brokerConfig struct {
	ClientPros string `opfield:"client_properties"`
}

// NewConsumer creates a new Kafka consumer and subscribes to the given topics.
func NewConsumer(ks *vault.KeyStoreConfig, topics []string) (*kafka.Consumer, error) {
	var brokerCfg brokerConfig
	if err := ks.Vault.LoadStructFromItemByTitle(&brokerCfg, vault.CoreBrokerConfigItem, ks.Env.Vault); err != nil {
		log.Printf("failed to load from key vault: %v", err)
		return nil, fmt.Errorf("failed to load from key vault: %v", err)
	}
	
	// create a new consumer
	consumer, err := kafka.NewConsumer(readFromOptsString(brokerCfg.ClientPros))
	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}
	
	// subscribe to topics
	go func(c *kafka.Consumer, topics []string) {
		if err = c.SubscribeTopics(topics, nil); err != nil {
			log.Fatalf("Failed to subscribe to topics %+v: %s\n", topics, err)
		}
	}(consumer, topics)
	
	return consumer, nil
}

// ConsumeMessages consumes messages from Kafka and calls the handler function for each message.
func ConsumeMessages(ctx context.Context, c *kafka.Consumer, msgChan chan<- *kafka.Message) {
	log.Println("Consuming messages...")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.ReadMessage(-1)
			if err != nil {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}
			msgChan <- msg
		}
	}
}

func readFromOptsString(opts string) *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": opts,
		"group.id":          "group",
		"auto.offset.reset": "earliest",
	}
}
