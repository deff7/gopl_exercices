package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC = -273.15
	FreezingC     = 0
	BolinigC      = 100
)

const (
	AbsoluteZeroK = 0
	FreezingK     = 273.15
	BoilingK      = 373.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
