package common

import (
	"context"
	"fmt"
	"github.com/egapay-core/shared/cache"
	pb "github.com/egapay-core/shared/grpc/proto/core/eganow/api/payment"
	"github.com/egapay-core/shared/vault"
	"log"
	"sync"
)

const (
	payPartnerCacheKeyPrefix = "pay.partner.svc"

	serviceIdsExpiration = 0 // never expire
)

// ISharedRepository is an interface for a repository that provides shared functionality
// across the application
type ISharedRepository interface {
	// GetConnectionForClient returns the connection string for a pay partner grpc client
	GetConnectionForClient(context.Context, *pb.NameEnquiryRequest) (string, error)
}

type sharedRepository struct {
	ks  *vault.KeyStoreConfig
	rdb cache.ICache
	ISharedRepository
}

func NewSharedRepository(ks *vault.KeyStoreConfig) ISharedRepository {
	// create cache instance
	rdb := cache.NewRdbConfigurator(ks).NewConn(cache.ServiceConnectionCacheIndex)
	cacheInstance := cache.NewCache(rdb)

	// load data from vault into cache
	loadDataFromVaultIntoCache(ks, cacheInstance)

	return &sharedRepository{ks: ks, rdb: cacheInstance}
}

func (r *sharedRepository) GetConnectionForClient(ctx context.Context, req *pb.NameEnquiryRequest) (string, error) {
	// load hash from cache
	hashed, err := r.rdb.GetAllFromHash(ctx, fmt.Sprintf("%s:%s", payPartnerCacheKeyPrefix, req.GetPayPartnerServiceId()))
	if err != nil {
		log.Printf("failed to get hash from cache: %v", err)
		return "", err
	}
	var config partnerServiceConfig
	if err = mapToStruct(hashed, &config); err != nil {
		log.Printf("failed to map hash to struct: %v", err)
		return "", err
	}

	// get connection string
	return fmt.Sprintf("%s.%s:%s", config.ServiceAddr, r.ks.PayPartnerServices.Host, r.ks.PayPartnerServices.Port), nil
}

func loadDataFromVaultIntoCache(ks *vault.KeyStoreConfig, instance cache.ICache) {
	var config payPartnerServiceIDFromVault

	// load data from vault
	var vaultWg sync.WaitGroup
	vaultWg.Add(4)

	var mtnVault mtnVaultConfig
	go func(wg *sync.WaitGroup, ks *vault.KeyStoreConfig, config *mtnVaultConfig) {
		defer wg.Done()
		if err := ks.Vault.LoadStructFromItemByTitle(config, vault.CorePayPartnerServiceConfigItem, ks.Env.Vault); err != nil {
			log.Printf("failed to load from key vault: %v", err)
			return
		}
	}(&vaultWg, ks, &mtnVault)

	var airtelVault airtelVaultConfig
	go func(wg *sync.WaitGroup, ks *vault.KeyStoreConfig, config *airtelVaultConfig) {
		defer wg.Done()
		if err := ks.Vault.LoadStructFromItemByTitle(config, vault.CorePayPartnerServiceConfigItem, ks.Env.Vault); err != nil {
			log.Printf("failed to load from key vault: %v", err)
			return
		}
	}(&vaultWg, ks, &airtelVault)

	var telecelVault telecelVaultConfig
	go func(wg *sync.WaitGroup, ks *vault.KeyStoreConfig, config *telecelVaultConfig) {
		defer wg.Done()
		if err := ks.Vault.LoadStructFromItemByTitle(config, vault.CorePayPartnerServiceConfigItem, ks.Env.Vault); err != nil {
			log.Printf("failed to load from key vault: %v", err)
			return
		}
	}(&vaultWg, ks, &telecelVault)

	var ghipssVault ghipssVaultConfig
	go func(wg *sync.WaitGroup, ks *vault.KeyStoreConfig, config *ghipssVaultConfig) {
		defer wg.Done()
		if err := ks.Vault.LoadStructFromItemByTitle(config, vault.CorePayPartnerServiceConfigItem, ks.Env.Vault); err != nil {
			log.Printf("failed to load from key vault: %v", err)
			return
		}
	}(&vaultWg, ks, &ghipssVault)

	vaultWg.Wait()

	// update config
	config.MtnConfig = &mtnVault
	config.AirtelConfig = &airtelVault
	config.TelecelConfig = &telecelVault
	config.GhipssConfig = ghipssVault.Config
	connectionCfg := NewPayPartnerConnectionFromVault(&config)

	// load data from vault into cache
	var svcs []*partnerServiceConfig
	svcs = append(
		svcs,
		connectionCfg.GhipssConfig,

		connectionCfg.MtnConfig.EgapayMadApi,
		connectionCfg.MtnConfig.EgapayOpenApi,
		connectionCfg.AirtelConfig.Egapay,
		connectionCfg.TelecelConfig.Egapay,

		connectionCfg.MtnConfig.PospayMadApi,
		connectionCfg.MtnConfig.PospayOpenApi,
		connectionCfg.AirtelConfig.Pospay,
		connectionCfg.TelecelConfig.Pospay,
	)

	var wg sync.WaitGroup
	for _, svc := range svcs {
		wg.Add(1)
		go func(wg *sync.WaitGroup, instance cache.ICache, svc *partnerServiceConfig) {
			defer wg.Done()
			ctx := context.Background()
			storeServiceConfig(ctx, instance, svc)
		}(&wg, instance, svc)
	}

	wg.Wait()
	log.Printf("loaded %d pay partner service configs from vault into cache", len(svcs))
}

// storeServiceConfig - store the pay partner service config in cache
func storeServiceConfig(ctx context.Context, rdb cache.ICache, config *partnerServiceConfig) {
	if config == nil {
		log.Println("pay partner service config is nil...skip storing in cache")
		return
	}

	key := fmt.Sprintf("%s:%s", payPartnerCacheKeyPrefix, config.ServiceID)
	if err := rdb.StoreHash(ctx, key, structToMap(*config), serviceIdsExpiration); err != nil {
		log.Printf("failed to store data in cache: %v", err)
		return
	}
	log.Printf("stored pay partner service config in cache: %v", config.ServiceID)
}
