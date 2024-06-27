package broker

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/egapay-core/shared/vault"
	"log"
)

// NewProducer creates a new Kafka producer.
func NewProducer(ks *vault.KeyStoreConfig) (*kafka.Producer, error) {
	var brokerCfg brokerConfig
	if err := ks.Vault.LoadStructFromItemByTitle(&brokerCfg, vault.CoreBrokerConfigItem, ks.Env.Vault); err != nil {
		log.Printf("failed to load from key vault: %v", err)
		return nil, fmt.Errorf("failed to load from key vault: %v", err)
	}
	
	// read the client properties
	props := readFromOptsString(brokerCfg.ClientPros)
	
	// create a new producer
	p, err := kafka.NewProducer(&props)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	return p, nil
}

// ProduceMessage produces a message to the given topic.
func ProduceMessage(p *kafka.Producer, topic string, key, value []byte) error {
	deliveryChan := make(chan kafka.Event)
	if err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          value,
	}, deliveryChan); err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %w", m.TopicPartition.Error)
	}
	return nil
}
