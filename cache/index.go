package cache

const (
	// SetupCacheIndex - index for setup cache (setup data like countries)
	SetupCacheIndex int = iota

	// CustomerCacheIndex - index for customer cache (customer profile)
	CustomerCacheIndex

	// CredentialsCacheIndex - index for credentials cache (credentials data like email configurations)
	CredentialsCacheIndex

	// TokenCacheIndex - index for token cache (token data like JWT)
	TokenCacheIndex

	// PaymentTypeCacheIndex - index for payment type cache (payment type data like card, bank etc. for source & beneficiary a/cs)
	PaymentTypeCacheIndex

	// TransactionCacheIndex - index for transaction cache (transaction data like transaction status)
	TransactionCacheIndex

	// ServiceConnectionCacheIndex - index for service connection cache (service connection data like pay partner service id)
	ServiceConnectionCacheIndex
)
