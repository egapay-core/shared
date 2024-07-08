package money_transfer

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/egapay-core/shared/grpc/client"
	cpb "github.com/egapay-core/shared/grpc/proto/core/eganow/api/payment"
	"github.com/egapay-core/shared/modules/common"
	"github.com/egapay-core/shared/vault"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StatusQuerySource int

const (
	FromCache StatusQuerySource = iota
	FromElasticsearch
	FromDatabase // @fixme: to be discussed later
)

type IMoneyTransferRepository interface {
	GetAccountHolder(context.Context, *cpb.PaymentNameEnquiryRequest) (*cpb.PaymentStringValue, error)
	CollectMoney(context.Context, *cpb.PaymentMoneyTransferRequest) (*cpb.PaymentMoneyTransferResponse, error)
	PayoutMoney(context.Context, *cpb.PaymentMoneyTransferRequest) (*cpb.PaymentMoneyTransferResponse, error)
	CheckTransactionStatus(context.Context, StatusQuerySource, *cpb.PaymentStatusQueryRequest) (string, error)
}

type moneyTransferRepository struct {
	ks                 *vault.KeyStoreConfig
	sharedRepo         common.ISharedRepository
	collectionProducer *kafka.Producer
	payoutProducer     *kafka.Producer
}

func NewMoneyTransferRepository(ks *vault.KeyStoreConfig, collectionProducer, payoutProducer *kafka.Producer) IMoneyTransferRepository {
	return &moneyTransferRepository{
		ks:                 ks,
		collectionProducer: collectionProducer,
		payoutProducer:     payoutProducer,
		sharedRepo:         common.NewSharedRepository(ks),
	}
}

func (r *moneyTransferRepository) GetAccountHolder(ctx context.Context, req *cpb.PaymentNameEnquiryRequest) (*cpb.PaymentStringValue, error) {
	addr, err := r.sharedRepo.GetConnectionForNameEnquiryClient(ctx, req)
	if err != nil {
		return nil, err
	}
	
	conn, err := client.CreateSecureGrpcConnection(context.Background(), addr)
	if err != nil {
		return nil, err
	}
	
	neqClient := cpb.NewCoreNameEnquirySvcClient(conn)
	return neqClient.GetAccountHolderName(ctx, req)
}

func (*moneyTransferRepository) CollectMoney(ctx context.Context, req *cpb.PaymentMoneyTransferRequest) (*cpb.PaymentMoneyTransferResponse, error) {
	// @todo - implement this method
	
	return nil, status.Error(codes.Unimplemented, "method not implemented yet")
}

func (*moneyTransferRepository) PayoutMoney(ctx context.Context, req *cpb.PaymentMoneyTransferRequest) (*cpb.PaymentMoneyTransferResponse, error) {
	// @todo - implement this method
	
	return nil, status.Error(codes.Unimplemented, "method not implemented yet")
}

func (*moneyTransferRepository) CheckTransactionStatus(context.Context, StatusQuerySource, *cpb.PaymentStatusQueryRequest) (string, error) {
	// @todo - implement this method
	
	return "", status.Error(codes.Unimplemented, "method not implemented yet")
}
