package currency

import (
	"testing"
)

/*
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
*/
func TestCurrencyValid(t *testing.T) {
	for _, tt := range []struct {
		c    Currency
		fail bool
	}{
		{c: JPY},
		{c: CNY},
		{c: HKD},
		{c: KRW},
		{c: EUR},
		{c: USD},
		{c: GBP},
		{c: CAD},
		{c: AUD},
		{c: MXN},
		{c: Currency("test"), fail: true},
		{c: Currency("1234"), fail: true},
		{c: Currency("xbt"), fail: true},
	}{
		baseValid := Valid(tt.c)
		cValid := tt.c.Valid()

		if baseValid != cValid {
			t.Error("both Valid() and c.Valid() should match.")
		}

		if baseValid != !tt.fail {
			t.Error("wrong value for Valid()")
		}

		if cValid != !tt.fail {
			t.Error("wrong value for c.Valid()")
		}
	}
}
