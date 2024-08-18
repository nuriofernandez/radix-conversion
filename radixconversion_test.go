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

func TestHex37E(t *testing.T) {
	hex := number.Number{
		Value: "37E",
		Base:  bases.Hexadecimal,
	}
	decimal := hex.ConvertTo(bases.Decimal)
	if decimal.Value != "894" {
		t.Fatalf("Incorrect result, expected '894' but received '%s'", decimal.Value)
		return
	}
}

func TestBinary132(t *testing.T) {
	decimal := number.Number{
		Value: "132",
		Base:  bases.Decimal,
	}
	bin := decimal.ConvertTo(bases.Binary)
	if bin.Value != "10000100" {
		t.Fatalf("Incorrect result, expected '10000100' but received '%s'", bin.Value)
		return
	}
}
