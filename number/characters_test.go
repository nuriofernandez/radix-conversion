package number

import (
	"github.com/nuriofernandez/radix-conversion/base"
	"testing"
)

func TestCharactersB10(t *testing.T) {
	number := Number{
		"10",
		base.Base{},
	}

	characters := number.Characters()
	if characters[0] != '1' {
		t.Error("Expected '1', got ", characters[0])
	}
	if characters[1] != '0' {
		t.Error("Expected '0', got ", characters[1])
	}
}

func TestCharactersB16(t *testing.T) {
	number := Number{
		"F3",
		base.Base{},
	}

	characters := number.Characters()
	if characters[0] != 'F' {
		t.Error("Expected 'F', got ", characters[0])
	}
	if characters[1] != '3' {
		t.Error("Expected '3', got ", characters[1])
	}
}
