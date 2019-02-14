// Package currency describes valid currencies available to the entire application.
// It also exposes a service handler that can be converted into a binary to serve this as a HTTP service.
package currency

import (
	"fmt"

	"github.com/monzo/typhon"
)

// currencies is a map of currency to description.
var currencies = map[Currency]string{
	JPY: "Japanese Yen (¥)",
	CNY: "Chinese Yuan (RMB/¥)",
	HKD: "Hong Kong Dollars (HK$)",
	KRW: "South Korean Won (₩)",
	EUR: "Euro (€)",
	USD: "US Dollars ($)",
	GBP: "Pound Sterling (£)",
	CAD: "Canadian Dollars (CA$)",
	AUD: "Australian Dollars (AU$)",
	MXN: "Mexican Pesos ($)",
}

// Currency is a simple string type that adds information to a currency.
type Currency string

const (
	// JPY is Japanese Yen (¥)
	JPY Currency = "jpy"

	// CNY is Chinese Yuan (RMB/¥)
	CNY Currency = "cny"

	// HKD is Hong Kong Dollar (HK$)
	HKD Currency = "hkd"

	// KRW is South Korean Won (₩)
	KRW Currency = "krw"

	// EUR is Euro (€)
	EUR Currency = "eur"

	// USD is US Dollars ($)
	USD Currency = "usd"

	// GBP is Pound Sterling (£)
	GBP Currency = "gbp"

	// CAD is Canadian Dollars (CA$)
	CAD Currency = "cad"

	// AUD is Australian Dollars (AU$)
	AUD Currency = "aud"

	// MXN is Mexican Pesos ($)
	MXN Currency = "mxn"
)

// Describe describes a currency, giving its full name and symbol.
func (c Currency) Describe() string {
	if d, ok := currencies[c]; ok {
		return d
	}

	return fmt.Sprintf("Unknown currency: %s", c)
}

// Valid checks if a currency given is accepted.
func (c Currency) Valid() bool {
	return Valid(c)
}

// Valid checks if a given currency is accepted.
func Valid(c Currency) bool {
	switch c {
	case JPY, CNY, HKD, KRW, EUR, USD, GBP, CAD, AUD, MXN:
		return true
	}
	return false
}

// Service executes a request and returns a response, handling all boilerplate.
func Service(req typhon.Request) typhon.Response {
	return req.Response(currencies)
}

