package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/egapay-core/shared/cache"
	"github.com/egapay-core/shared/database/datasource"
	"github.com/egapay-core/shared/database/entities"
	"log"
)

const (
	cacheKeyOperatingCountries = "countries.operating"
	cacheKeySignupCountries    = "countries.signup"
)

type ICountryRepository interface {
	GetOperatingCountries(ctx context.Context) ([]*entities.CountryEntity, error)
	GetSignupCountries(ctx context.Context) ([]*entities.CountryEntity, error)
}

type countryRepository struct {
	db    datasource.ICountryDataSource
	cache cache.ICache
}

func NewCountryRepository(cache cache.ICache, db datasource.ICountryDataSource) ICountryRepository {
	return &countryRepository{db: db, cache: cache}
}

func (r *countryRepository) GetOperatingCountries(ctx context.Context) ([]*entities.CountryEntity, error) {
	// get countries from cache
	countries, err := r.cache.GetList(ctx, cacheKeyOperatingCountries)
	if err != nil {
		// get countries from db
		countries, err := r.db.DGetOperatingCountries(ctx)
		if err != nil {
			return nil, err
		}

		var tempCountries []*interface{}
		for _, country := range countries {
			var temp interface{} = country
			tempCountries = append(tempCountries, &temp)
		}

		// store countries in cache
		if err = r.cache.StoreList(ctx, cacheKeyOperatingCountries, tempCountries, 0); err != nil {
			log.Printf("failed to store countries in cache: %v", err)
			return nil, err
		}

		return countries, nil
	}

	// convert countries from []*interface to []*entities.CountryEntity
	var tempCountries []*entities.CountryEntity
	for _, country := range countries {
		var tempCountry entities.CountryEntity
		countryMap, ok := (*country).(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to cast country to map[string]interface{}: %v", err)
		}
		countryBytes, err := json.Marshal(countryMap)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal country map to bytes: %v", err)
		}
		if err = json.Unmarshal(countryBytes, &tempCountry); err != nil {
			return nil, fmt.Errorf("failed to unmarshal country bytes to *entities.CountryEntity: %v", err)
		}
		tempCountries = append(tempCountries, &tempCountry)
	}

	return tempCountries, nil
}

func (r *countryRepository) GetSignupCountries(ctx context.Context) ([]*entities.CountryEntity, error) {
	// get countries from cache
	countries, err := r.cache.GetList(ctx, cacheKeySignupCountries)
	if err != nil {
		// get countries from db
		countries, err := r.db.DGetSignupCountries(ctx)
		if err != nil {
			return nil, err
		}

		var tempCountries []*interface{}
		for _, country := range countries {
			var temp interface{} = country
			tempCountries = append(tempCountries, &temp)
		}

		// store countries in cache
		if err = r.cache.StoreList(ctx, cacheKeySignupCountries, tempCountries, 0); err != nil {
			log.Printf("failed to store countries in cache: %v", err)
			return nil, err
		}

		return countries, nil
	}

	// convert countries from []*interface to []*entities.CountryEntity
	var tempCountries []*entities.CountryEntity
	for _, country := range countries {
		var tempCountry entities.CountryEntity
		countryMap, ok := (*country).(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to cast country to map[string]interface{}: %v", err)
		}
		countryBytes, err := json.Marshal(countryMap)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal country map to bytes: %v", err)
		}
		if err = json.Unmarshal(countryBytes, &tempCountry); err != nil {
			return nil, fmt.Errorf("failed to unmarshal country bytes to *entities.CountryEntity: %v", err)
		}
		tempCountries = append(tempCountries, &tempCountry)
	}

	return tempCountries, nil
}
