/* Summary comments that document the package...in this case, this package provides
datatypes for Celsius and Fahrenheit, and utilities to convert b/w each other*/
package tempconv

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)
