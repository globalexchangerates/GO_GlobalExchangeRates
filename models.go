package globalexchangerates

// Currency represents a currency in the exchange rates system.
type Currency struct {
	// Code is the ISO code of the currency.
	Code string `json:"code"`
	
	// Name is the display name of the currency.
	Name string `json:"name"`
	
	// NumericCode is the numeric ISO code of the currency.
	NumericCode string `json:"numericCode"`
	
	// Obsolete indicates whether the currency is obsolete.
	Obsolete bool `json:"obsolete"`
}

// Provider represents a data provider for exchange rates.
type Provider struct {
	// Code is the provider code.
	Code string `json:"code"`
	
	// Description is the provider description.
	Description string `json:"description"`
	
	// Country is the country name where the provider is based.
	Country string `json:"country"`
	
	// ReferenceCurrency is the reference currency code used by the provider.
	ReferenceCurrency string `json:"referenceCurrency"`
	
	// TimeSeries indicates whether the provider supports time series data.
	TimeSeries bool `json:"timeSeries"`
	
	// Monthly indicates whether the provider provides monthly data.
	Monthly bool `json:"monthly"`
	
	// CountryCode is the country code where the provider is based.
	CountryCode string `json:"countryCode"`
}

// ExchangeRateResponse represents an exchange rate response from the API.
type ExchangeRateResponse struct {
	// Provider is the provider code.
	Provider string `json:"provider"`
	
	// Date is the date for the exchange rates in YYYY-MM-DD format.
	Date CustomTime `json:"date"`
	
	// Base is the base currency code.
	Base string `json:"base"`
	
	// ExchangeRates is the dictionary where the key is the currency code and the value is the exchange rate.
	ExchangeRates map[string]float64 `json:"exchangeRates"`
}

// ConversionResponse represents a currency conversion response from the API.
type ConversionResponse struct {
	// Provider is the provider code.
	Provider string `json:"provider"`
	
	// Date is the date for the conversion rates in YYYY-MM-DD format.
	Date CustomTime `json:"date"`
	
	// Base is the base currency code.
	Base string `json:"base"`
	
	// Amount is the amount to convert.
	Amount float64 `json:"amount"`
	
	// Conversions is the dictionary where the key is the currency code and the value is the converted amount.
	Conversions map[string]float64 `json:"conversions"`
}

// ErrorResponse represents an error response from the API.
type ErrorResponse struct {
	// Message is the error message.
	Message string `json:"message"`
	
	// ErrorCode is the error code.
	ErrorCode int `json:"errorCode"`
}
