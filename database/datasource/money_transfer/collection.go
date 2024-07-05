package money_transfer

import "github.com/jackc/pgx/v5/pgxpool"

type ICollectionDataSource interface {
	CollectMoney() error
}

type collectionDataSource struct {
	crdb *pgxpool.Pool
	ICollectionDataSource
}

func NewCollectionDataSource(crdb *pgxpool.Pool) ICollectionDataSource {
	return &collectionDataSource{crdb: crdb}
}
