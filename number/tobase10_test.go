package number

import (
	"binary-conversion/base"
	"testing"
)

var base10 = base.Base{
	Name: "Base 10",
	Characters: map[int]rune{
		0: '0',
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
	},
	Radix: 10,
}

var hexadecimal = base.Base{
	Name: "Hexadecimal",
	Characters: map[int]rune{
		0:  '0',
		1:  '1',
		2:  '2',
		3:  '3',
		4:  '4',
		5:  '5',
		6:  '6',
		7:  '7',
		8:  '8',
		9:  '9',
		10: 'A',
		11: 'B',
		12: 'C',
		13: 'D',
		14: 'E',
		15: 'F',
	},
	Radix: 16,
}

func TestDecimal10(t *testing.T) {
	number := Number{
		Value: "10",
		Base:  base10,
	}

	toBase10 := number.ToBase10()
	if toBase10 != 10 {
		t.Errorf("Expected '10' but got '%f'", toBase10)
	}
}

func TestDecimal9714(t *testing.T) {
	number := Number{
		Value: "9714",
		Base:  base10,
	}

	toBase10 := number.ToBase10()
	if toBase10 != 9714 {
		t.Errorf("Expected '9714' but got '%f'", toBase10)
	}
}

func TestHex37E(t *testing.T) {
	number := Number{
		Value: "37E",
		Base:  hexadecimal,
	}

	toBase10 := number.ToBase10()
	if toBase10 != 894 {
		t.Errorf("Expected '894' but got '%f'", toBase10)
	}
}
