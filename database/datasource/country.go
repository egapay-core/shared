package datasource

import (
	"context"
	"github.com/egapay-core/shared/database/entities"
	"github.com/egapay-core/shared/database/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ICountryDataSource - interface for country data source
type ICountryDataSource interface {
	DGetOperatingCountries(context.Context) ([]*entities.CountryEntity, error)
	DGetSignupCountries(context.Context) ([]*entities.CountryEntity, error)
}

type countryDataSource struct {
	crdb *pgxpool.Pool
	ICountryDataSource
}

func NewCountryDataSource(crdb *pgxpool.Pool) ICountryDataSource {
	return &countryDataSource{crdb: crdb}
}

func (r *countryDataSource) DGetOperatingCountries(ctx context.Context) ([]*entities.CountryEntity, error) {
	countries := make([]*entities.CountryEntity, 0)
	rows, err := r.crdb.Query(ctx, `CALL getoperatingcountries()`)

	for rows.Next() {
		country := new(entities.CountryEntity)
		if err = rows.Scan(&country.CountryCode, &country.DialCode, &country.Name, &country.IsoCode,
			&country.Status, &country.AllowForCustomerOnboarding, &country.CurrencyName,
			&country.FlagUrl, &country.MinMobileNumberLength, &country.MaxMobileNumberLength); err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}

	return countries, nil
}

func (r *countryDataSource) DGetSignupCountries(ctx context.Context) ([]*entities.CountryEntity, error) {
	countries := make([]*entities.CountryEntity, 0)
	rows, err := r.crdb.Query(ctx, `
SELECT countrycode, countrydialcode, countryname, currencyiso,countrystatusactiveorclosed,
       allowforcustomeronboardingyesno, currencyname, countryflagurl, mobilenominlen, mobilenomaxlen
FROM setupcountry
WHERE blacklistedyesno = $1 and countrystatusactiveorclosed = $2
`, string(utils.No), string(utils.Active))

	for rows.Next() {
		country := new(entities.CountryEntity)
		if err = rows.Scan(&country.CountryCode, &country.DialCode, &country.Name, &country.IsoCode,
			&country.Status, &country.AllowForCustomerOnboarding, &country.CurrencyName,
			&country.FlagUrl, &country.MinMobileNumberLength, &country.MaxMobileNumberLength); err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}

	return countries, nil
}
