package money_transfer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/egapay-core/shared/vault"
)

type IMoneyTransferRepository interface {
	// @todo - add methods here
}

type moneyTransferRepository struct {
	ks                 *vault.KeyStoreConfig
	collectionProducer *kafka.Producer
	payoutProducer     *kafka.Producer
}

func NewMoneyTransferRepository(ks *vault.KeyStoreConfig, collectionProducer, payoutProducer *kafka.Producer) IMoneyTransferRepository {
	return &moneyTransferRepository{
		ks:                 ks,
		collectionProducer: collectionProducer,
		payoutProducer:     payoutProducer,
	}
}
