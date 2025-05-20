package globalexchangerates

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// GetCurrenciesOptions represents the options for the GetCurrencies method.
type GetCurrenciesOptions struct {
	// Codes is the optional filter to retrieve specific currency codes.
	Codes []string
}

// GetCurrencies gets a list of supported currencies.
func (c *Client) GetCurrencies(ctx context.Context, options *GetCurrenciesOptions) ([]Currency, error) {
	params := url.Values{}
	
	if options != nil && len(options.Codes) > 0 {
		params.Set("code", strings.Join(options.Codes, ","))
	}
	
	var currencies []Currency
	err := c.sendRequest(ctx, "/currencies", params, &currencies)
	if err != nil {
		return nil, err
	}
	
	return currencies, nil
}

// GetLatestOptions represents the options for the GetLatest method.
type GetLatestOptions struct {
	// Provider is the optional provider code.
	Provider string
	
	// Currencies is the optional list of currencies to include in the response.
	Currencies []string
	
	// BaseCurrency is the optional base currency for the rates.
	BaseCurrency string
}

// GetLatest gets the latest exchange rates.
func (c *Client) GetLatest(ctx context.Context, options *GetLatestOptions) (*ExchangeRateResponse, error) {
	params := url.Values{}
	
	if options != nil {
		if options.Provider != "" {
			params.Set("provider", options.Provider)
		}
		
		if len(options.Currencies) > 0 {
			params.Set("currencies", strings.Join(options.Currencies, ","))
		}
		
		if options.BaseCurrency != "" {
			params.Set("base", options.BaseCurrency)
		}
	}
	
	rates := &ExchangeRateResponse{}
	err := c.sendRequest(ctx, "/latest", params, rates)
	if err != nil {
		return nil, err
	}
	
	return rates, nil
}

// GetProvidersOptions represents the options for the GetProviders method.
type GetProvidersOptions struct {
	// Codes is the optional filter to retrieve specific provider codes.
	Codes []string
	
	// CountryCode is the optional filter to retrieve providers from a specific country.
	CountryCode string
}

// GetProviders gets a list of supported providers.
func (c *Client) GetProviders(ctx context.Context, options *GetProvidersOptions) ([]Provider, error) {
	params := url.Values{}
	
	if options != nil {
		if len(options.Codes) > 0 {
			params.Set("code", strings.Join(options.Codes, ","))
		}
		
		if options.CountryCode != "" {
			params.Set("countryCode", options.CountryCode)
		}
	}
	
	var providers []Provider
	err := c.sendRequest(ctx, "/providers", params, &providers)
	if err != nil {
		return nil, err
	}
	
	return providers, nil
}

// GetHistoricalOptions represents the options for the GetHistorical method.
type GetHistoricalOptions struct {
	// Latest is the optional flag to get the latest rates for the specified date.
	Latest bool
	
	// Provider is the optional provider code.
	Provider string
	
	// Currencies is the optional list of currencies to include in the response.
	Currencies []string
	
	// BaseCurrency is the optional base currency for the rates.
	BaseCurrency string
}

// GetHistorical gets historical exchange rates for a specific date.
func (c *Client) GetHistorical(ctx context.Context, date time.Time, options *GetHistoricalOptions) (*ExchangeRateResponse, error) {
	params := url.Values{}
	
	// Set the date parameter in YYYY-MM-DD format
	params.Set("date", date.Format("2006-01-02"))
	
	if options != nil {
		if options.Latest {
			params.Set("latest", "true")
		}
		
		if options.Provider != "" {
			params.Set("provider", options.Provider)
		}
		
		if len(options.Currencies) > 0 {
			params.Set("currencies", strings.Join(options.Currencies, ","))
		}
		
		if options.BaseCurrency != "" {
			params.Set("base", options.BaseCurrency)
		}
	}
	
	rates := &ExchangeRateResponse{}
	err := c.sendRequest(ctx, "/historical", params, rates)
	if err != nil {
		return nil, err
	}
	
	return rates, nil
}

// ConvertOptions represents the options for the Convert method.
type ConvertOptions struct {
	// BaseCurrency is the optional source currency code.
	BaseCurrency string
	
	// ToCurrencies is the optional target currency codes.
	ToCurrencies []string
	
	// Provider is the optional provider code.
	Provider string
	
	// Date is the optional date for historical conversions.
	Date *time.Time
}

// Convert converts an amount from one currency to others.
func (c *Client) Convert(ctx context.Context, amount float64, options *ConvertOptions) (*ConversionResponse, error) {
	params := url.Values{}
	
	// Add amount parameter
	params.Set("amount", formatFloat(amount))
	
	if options != nil {
		if options.BaseCurrency != "" {
			params.Set("base", options.BaseCurrency)
		}
		
		if len(options.ToCurrencies) > 0 {
			params.Set("to", strings.Join(options.ToCurrencies, ","))
		}
		
		if options.Provider != "" {
			params.Set("provider", options.Provider)
		}
		
		if options.Date != nil {
			params.Set("date", options.Date.Format("2006-01-02"))
		}
	}
	
	conversion := &ConversionResponse{}
	err := c.sendRequest(ctx, "/convert", params, conversion)
	if err != nil {
		return nil, err
	}
	
	return conversion, nil
}

// Helper function to format float values
func formatFloat(f float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.6f", f), "0"), ".")
}
