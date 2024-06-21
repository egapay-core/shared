package datasource

import (
	"context"
	"github.com/egapay-core/shared/database/entities"
)

// ICountryDataSource - interface for country data source
type ICountryDataSource interface {
	GetOperatingCountries(context.Context) ([]*entities.CountryEntity, error)
	GetSignupCountries(context.Context) ([]*entities.CountryEntity, error)
}
