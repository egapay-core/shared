package money_transfer

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	cpb "github.com/egapay-core/shared/grpc/proto/core/eganow/api/payment"
	"github.com/egapay-core/shared/vault"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type IMoneyTransferRepository interface {
	GetAccountHolder(context.Context, *cpb.PaymentNameEnquiryRequest) (*cpb.PaymentNameEnquiryResponse, error)
	CollectMoney(context.Context, *cpb.PaymentMoneyTransferRequest) (*cpb.PaymentMoneyTransferResponse, error)
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

func (*moneyTransferRepository) GetAccountHolder(ctx context.Context, req *cpb.PaymentNameEnquiryRequest) (*cpb.PaymentNameEnquiryResponse, error) {
	// @todo - implement this method
	addr := "192.168.1.119:50051"
	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	//client.CreateSecureGrpcConnection(context.Background(), "192.168.1.119:50051")
	client := cpb.NewCoreNameEnquirySvcClient(conn)
	return client.GetAccountHolderName(ctx, req)
}

func (*moneyTransferRepository) CollectMoney(ctx context.Context, req *cpb.PaymentMoneyTransferRequest) (*cpb.PaymentMoneyTransferResponse, error) {
	// @todo - implement this method

	return nil, status.Error(codes.Unimplemented, "method not implemented yet")
}