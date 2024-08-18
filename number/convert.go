package number

import (
	"binary-conversion/base"
	"slices"
)

func (n *Number) ConvertTo(base base.Base) Number {
	radix := float64(base.Radix)

	results := make([]int, 0)
	next := n.ToBase10()
	for {
		if next <= 0 {
			break
		}
		var result = next / radix

		next = float64(int(result))

		rest := result - next
		index := int(rest * radix)

		results = append(results, index)
	}

	// Set reverse order
	slices.Reverse(results)

	newValue := ""
	for _, charIndex := range results {
		destinationChar := base.ConvertFromIndex(charIndex)
		newValue += string(destinationChar)
	}

	return Number{
		Value: newValue,
		Base:  base,
	}
}
