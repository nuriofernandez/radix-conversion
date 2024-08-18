package base

func (b *Base) ConvertFromIndex(number int) rune {
	return b.Characters[number]
}

func (b *Base) ConvertToIndex(value rune) int {
	for i, r := range b.Characters {
		if r == value {
			return i
		}
	}
	panic("unable to convert to index")
}
