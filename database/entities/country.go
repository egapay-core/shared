package entities

import "github.com/egapay-core/shared/database/utils"

// CountryEntity - Model for country table
type CountryEntity struct {
	CountryCode                string                 `json:"countrycode" db:"countrycode"`
	DialCode                   string                 `json:"countrydialcode" db:"countrydialcode"`
	Name                       string                 `json:"countryname" db:"countryname"`
	IsoCode                    string                 `json:"currencyiso" db:"currencyiso"`
	Status                     utils.ActiveClosedEnum `json:"countrystatusactiveorclosed" db:"countrystatusactiveorclosed"`
	AllowForCustomerOnboarding utils.YesNoEnum        `json:"allowforcustomeronboardingyesno" db:"allowforcustomeronboardingyesno"`
	CurrencyName               string                 `json:"currencyname" db:"currencyname"`
	FlagUrl                    string                 `json:"countryflagurl" db:"countryflagurl"`
	MinMobileNumberLength      int                    `json:"mobilenominlen" db:"mobilenominlen"`
	MaxMobileNumberLength      int                    `json:"mobilenomaxlen" db:"mobilenomaxlen"`
}
