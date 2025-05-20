// Package globalexchangerates provides a client for interacting with the Global Exchange Rates API.
//
// The Global Exchange Rates API (https://globalexchangerates.org/) provides a simple and reliable way to download
// official exchange rates from central banks and tax authorities (providers) worldwide.
//
// This package provides a set of methods for retrieving exchange rates, currency and provider information,
// and for converting amounts between currencies.
//
// # Getting Started
//
// To use this package, you need to sign up to the Developer API Portal (https://dev.globalexchangerates.org/Account/Signup)
// and obtain an API key. Then, create a new client with your API key:
//
//	client := globalexchangerates.NewClient("your_api_key")
//
// You can then call methods to interact with the API, such as retrieving exchange rates:
//
//	rates, err := client.GetLatest(context.Background(), nil)
//
// # Error Handling
//
// Methods in this package return errors if the API request fails. You can check if an error is an APIError:
//
//	if apiErr, ok := globalexchangerates.IsAPIError(err); ok {
//	    // Handle API error
//	    fmt.Println("API Error:", apiErr.StatusCode, apiErr.ErrorCode, apiErr.Message)
//	} else {
//	    // Handle other error
//	    fmt.Println("Error:", err)
//	}
//
// # Advanced Usage
//
// For more advanced use cases, you can provide options to the API methods:
//
//	options := &globalexchangerates.GetLatestOptions{
//	    Provider:     "ECB",
//	    BaseCurrency: "USD",
//	    Currencies:   []string{"EUR", "GBP", "JPY"},
//	}
//	rates, err := client.GetLatest(context.Background(), options)
//
// You can also customize the HTTP client:
//
//	client := globalexchangerates.NewClient("your_api_key",
//	    globalexchangerates.WithHTTPClient(&http.Client{
//	        Timeout: 10 * time.Second,
//	    }),
//	)
package globalexchangerates
