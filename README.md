# Global Exchange Rates Client

This is a Go library designed to interact with the [Global Exchange Rates API](https://globalexchangerates.org/).

It provides a set of methods for retrieving exchange rates, currency and providers information, and convert amounts between currencies.

## Global Exchange Rates API

The **Global Exchange Rates API** provides a simple and reliable way to download **official exchange rates** from central banks and tax authorities (*providers*) worldwide. 

You can check the updated list from the [official website](https://www.globalexchangerates.org/global-coverage/).

**Daily exchange rates** of all the available currencies are also provided, calculated using a proprietary algorithm blending the official data from central banks and the market rate from commercial institutions.

Ideal for accounting, CRM and ERP systems, business intelligence, auditing, tax compliance and reporting.

## Getting Started

1. Sign up to the [Developer API Portal](https://dev.globalexchangerates.org/Account/Signup) and start the 30 day free trial.
2. Get your API key.
3. Install the package:
   ```bash
   go get github.com/globalexchangerates/GO_GlobalExchangeRates 
   ```
4. Use the API.

## Usage

To use it, create an instance of `Client` with your API key:

```go
client := globalexchangerates.NewClient("your_api_key")
```

You can then call methods to interact with the API, such as retrieving exchange rates:

```go
// Latest Rates
rates, err := client.GetLatest(context.Background(), nil)
if err != nil {
    log.Fatal(err)
}
```

## License

This project is licensed under the MIT License.

## Full Documentation

The full API documentation is available at [doc.globalexchangerates.org](https://doc.globalexchangerates.org/).

## Contacts and support
[support@globalexchangerates.org](mailto:support@globalexchangerates.org).