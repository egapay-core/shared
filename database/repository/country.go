package repository

import (
	"context"
	"github.com/egapay-core/shared/database/datasource"
	"github.com/egapay-core/shared/database/entities"
	"github.com/egapay-core/shared/database/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type countryRepository struct {
	crdb *pgxpool.Conn
	datasource.ICountryDataSource
}

func NewCountryRepository(crdb *pgxpool.Conn) datasource.ICountryDataSource {
	return &countryRepository{crdb: crdb}
}

func (r *countryRepository) GetOperatingCountries(ctx context.Context) ([]*entities.CountryEntity, error) {
	countries := make([]*entities.CountryEntity, 0)
	rows, err := r.crdb.Query(ctx, `
SELECT countrycode, countrydialcode, countryname, currencyiso,countrystatusactiveorclosed,
       allowforcustomeronboardingyesno, currencyname, countryflagurl, mobilenominlen, mobilenomaxlen
FROM setupcountry
WHERE operatingcountryyesno = $1 and blacklistedyesno = $2 and countrystatusactiveorclosed = $3
`, string(utils.Yes), string(utils.No), string(utils.Active))
	
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

func (r *countryRepository) GetSignupCountries(ctx context.Context) ([]*entities.CountryEntity, error) {
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
