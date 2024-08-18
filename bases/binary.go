package bases

import (
	"github.com/nuriofernandez/radix-conversion/base"
)

var Binary = base.Base{
	Name: "Binary",
	Characters: map[int]rune{
		0: '0',
		1: '1',
	},
	Radix: 2,
}
