package main

import (
	"binary-conversion/base"
	"binary-conversion/number"
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

func TestDecimal17(t *testing.T) {
	decimal := number.Number{
		Value: "17",
		Base:  base10,
	}
	hex := decimal.ConvertTo(hexadecimal)
	if hex.Value != "11" {
		t.Fatalf("Incorrect result, expected '11' but received '%s'", hex.Value)
		return
	}
}

func TestDecimal405(t *testing.T) {
	decimal := number.Number{
		Value: "405",
		Base:  base10,
	}
	hex := decimal.ConvertTo(hexadecimal)
	if hex.Value != "195" {
		t.Fatalf("Incorrect result, expected '195' but received '%s'", hex.Value)
		return
	}
}
