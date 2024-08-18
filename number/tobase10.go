package number

import (
	"math"
	"slices"
)

func (n *Number) ToBase10() float64 {
	characters := n.Characters()

	// Reverse order (so pow(r, i) will work properly without the 'size-1-i' math)
	slices.Reverse(characters)

	// Perform the conversion
	result := float64(0)
	for i, character := range characters {
		base10 := n.Base.ConvertToIndex(character)
		result += float64(base10) * math.Pow(float64(n.Base.Radix), float64(i))
	}
	return result
}
