package radixconversion

import (
	"github.com/nuriofernandez/radix-conversion/bases"
	"github.com/nuriofernandez/radix-conversion/number"
	"testing"
)

func TestDecimal17(t *testing.T) {
	decimal := number.Number{
		Value: "17",
		Base:  bases.Decimal,
	}
	hex := decimal.ConvertTo(bases.Hexadecimal)
	if hex.Value != "11" {
		t.Fatalf("Incorrect result, expected '11' but received '%s'", hex.Value)
		return
	}
}

func TestDecimal405(t *testing.T) {
	decimal := number.Number{
		Value: "405",
		Base:  bases.Decimal,
	}
	hex := decimal.ConvertTo(bases.Hexadecimal)
	if hex.Value != "195" {
		t.Fatalf("Incorrect result, expected '195' but received '%s'", hex.Value)
		return
	}
}
