# Radix conversion

This package is a project to convert numbers from different base systems.

It allows to specify the radix of a base and define your custom character mapping.

# Decimal (b10) to Hexadecimal (b16) example:

### Definition of the base systems:

```go
var Decimal = base.Base{
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
```

```go

var Hexadecimal = base.Base{
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
```

### Operation

```go
decimal := number.Number{
    Value: "894",
    Base:  bases.Decimal,
}

hex := decimal.ConvertTo(bases.Hexadecimal)
fmt.Println("Result is: " + hex.Value)
```

#### Output:

> Result is: 37E
