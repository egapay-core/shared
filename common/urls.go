package common

import "reflect"

// PayPartnerServiceConnectionString - holds the pay partner service connection strings
type PayPartnerServiceConnectionString struct {
	MtnConfig     *mtnConfig            `json:"mtn"`
	TelecelConfig *generalPartnerConfig `json:"telecel"`
	AirtelConfig  *generalPartnerConfig `json:"airtel"`
	GhipssConfig  *partnerServiceConfig `json:"ghipss"`
}

// NewPayPartnerConnectionFromVault - creates a new pay partner connection from vault
func NewPayPartnerConnectionFromVault(cfg *payPartnerServiceIDFromVault) *PayPartnerServiceConnectionString {
	return &PayPartnerServiceConnectionString{
		MtnConfig: &mtnConfig{
			EgapayMadApi: &partnerServiceConfig{
				ServiceID:   cfg.MtnConfig.EgapayMadApi,
				ServiceAddr: "mtnegapaymadapi",
			},
			PospayMadApi: &partnerServiceConfig{
				ServiceID:   cfg.MtnConfig.PospayMadApi,
				ServiceAddr: "mtnpospaymadapi",
			},
			EgapayOpenApi: &partnerServiceConfig{
				ServiceID:   cfg.MtnConfig.EgapayOpenApi,
				ServiceAddr: "mtnegapayopenapi",
			},
			PospayOpenApi: &partnerServiceConfig{
				ServiceID:   cfg.MtnConfig.PospayOpenApi,
				ServiceAddr: "mtnpospayopenapi",
			},
		},
		TelecelConfig: &generalPartnerConfig{
			Egapay: &partnerServiceConfig{
				ServiceID:   cfg.TelecelConfig.Egapay,
				ServiceAddr: "telecelegapay",
			},
			Pospay: &partnerServiceConfig{
				ServiceID:   cfg.TelecelConfig.Pospay,
				ServiceAddr: "telecelpospay",
			},
		},
		AirtelConfig: &generalPartnerConfig{
			Egapay: &partnerServiceConfig{
				ServiceID:   cfg.AirtelConfig.Egapay,
				ServiceAddr: "airtelegapay",
			},
			Pospay: &partnerServiceConfig{
				ServiceID:   cfg.AirtelConfig.Pospay,
				ServiceAddr: "airtelpospay",
			},
		},
		GhipssConfig: &partnerServiceConfig{
			ServiceID:   cfg.GhipssConfig,
			ServiceAddr: "ghipss",
		},
	}
}

// partnerServiceConfig - holds the partner service configuration
type partnerServiceConfig struct {
	ServiceAddr string `json:"service_addr"`
	ServiceID   string `json:"service_id"`
}

// mtnConfig - holds the mtn pay partner configuration
type mtnConfig struct {
	EgapayMadApi  *partnerServiceConfig `json:"egapay_madapi"`
	PospayMadApi  *partnerServiceConfig `json:"pospay_madapi"`
	EgapayOpenApi *partnerServiceConfig `json:"egapay_openapi"`
	PospayOpenApi *partnerServiceConfig `json:"pospay_openapi"`
}

// generalPartnerConfig - holds the general partner configuration
type generalPartnerConfig struct {
	Egapay *partnerServiceConfig `json:"egapay"`
	Pospay *partnerServiceConfig `json:"pospay"`
}

type mtnVaultConfig struct {
	EgapayMadApi  string `opfield:"egapay_mtn_madapi_sid"`
	EgapayOpenApi string `opfield:"egapay_mtn_openapi_sid"`
	PospayMadApi  string `opfield:"pospay_mtn_madapi_sid"`
	PospayOpenApi string `opfield:"pospay_mtn_openapi_sid"`
}

type telecelVaultConfig struct {
	Egapay string `opfield:"egapay_telecel_sid"`
	Pospay string `opfield:"pospay_telecel_sid"`
}

type airtelVaultConfig struct {
	Egapay string `opfield:"egapay_airtel_sid"`
	Pospay string `opfield:"pospay_airtel_sid"`
}

type ghipssVaultConfig struct {
	Config string `opfield:"ghipss_sid"`
}

// payPartnerServiceIDFromVault - holds the pay partner service IDs from vault
type payPartnerServiceIDFromVault struct {
	MtnConfig     *mtnVaultConfig
	TelecelConfig *telecelVaultConfig
	AirtelConfig  *airtelVaultConfig
	GhipssConfig  string
}

// structToMap - converts a struct to a map
func structToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	value := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typ.Field(i)
		fieldName := fieldType.Name

		if field.Kind() == reflect.Ptr && !field.IsNil() {
			result[fieldName] = structToMap(field.Elem().Interface())
		} else {
			result[fieldName] = field.Interface()
		}
	}

	return result
}

// mapToStruct - converts a map to a struct
func mapToStruct(data map[string]interface{}, result interface{}) error {
	val := reflect.ValueOf(result).Elem()
	for k, v := range data {
		field := val.FieldByName(k)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(v).Convert(field.Type()))
		}
	}
	return nil
}
