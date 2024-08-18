package number

import (
	"github.com/nuriofernandez/radix-conversion/bases"
	"testing"
)

func TestDecimal10(t *testing.T) {
	number := Number{
		Value: "10",
		Base:  bases.Decimal,
	}

	toBase10 := number.ToBase10()
	if toBase10 != 10 {
		t.Errorf("Expected '10' but got '%f'", toBase10)
	}
}

func TestDecimal9714(t *testing.T) {
	number := Number{
		Value: "9714",
		Base:  bases.Decimal,
	}

	toBase10 := number.ToBase10()
	if toBase10 != 9714 {
		t.Errorf("Expected '9714' but got '%f'", toBase10)
	}
}

func TestHex37E(t *testing.T) {
	number := Number{
		Value: "37E",
		Base:  bases.Hexadecimal,
	}

	toBase10 := number.ToBase10()
	if toBase10 != 894 {
		t.Errorf("Expected '894' but got '%f'", toBase10)
	}
}
