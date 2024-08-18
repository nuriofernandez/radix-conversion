package main

import (
	"fmt"
	"github.com/nuriofernandez/radix-conversion/base"
	"github.com/nuriofernandez/radix-conversion/number"
)

func main() {
	fmt.Println("Conversion redix :)")

	base10 := base.Base{
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

	hexadecimal := base.Base{
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

	decimal := number.Number{
		Value: "3284",
		Base:  base10,
	}

	fmt.Printf(
		"Converting number '%s'@'%s' to '%s'\n",
		decimal.Value, decimal.Base.Name, hexadecimal.Name,
	)

	to := decimal.ConvertTo(hexadecimal)

	fmt.Println(to.Value)
}
