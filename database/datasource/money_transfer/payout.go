package money_transfer

import "github.com/jackc/pgx/v5/pgxpool"

type IPayoutDataSource interface {
	PayoutMoney() error
}

type payoutDataSource struct {
	crdb *pgxpool.Pool
	IPayoutDataSource
}

func NewPayoutDataSource(crdb *pgxpool.Pool) IPayoutDataSource {
	return &payoutDataSource{crdb: crdb}
}
