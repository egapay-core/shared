package cache

const (
	// CustomerCacheIndex - index for customer cache (customer profile)
	CustomerCacheIndex int = iota
	
	// SetupCacheIndex - index for setup cache (setup data like countries)
	SetupCacheIndex
	
	// CredentialsCacheIndex - index for credentials cache (credentials data like email configurations)
	CredentialsCacheIndex
	
	// PaymentTypeCacheIndex - index for payment type cache (payment type data like card, bank etc. for source & beneficiary a/cs)
	PaymentTypeCacheIndex
)
