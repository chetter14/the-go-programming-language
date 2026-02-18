package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroKelvin            = 0
)

func (c Celsius) String() string { return fmt.Sprintf("%g*C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g*F", f) }

func (k Kelvin) String() string { return fmt.Sprintf("%g*K", k) }

// func main() {
// fmt.Println(CToF(50))
// fmt.Println(FToC(82))

// fmt.Printf("%g\n", FreezingC + FToC(0))

// boilingF := CToF(BoilingC)
// fmt.Printf("%g\n", boilingF)
// // fmt.Printf("%g\n", boilingF - BoilingC)		/* compile error */

// fmt.Println(FreezingC)
// c := FToC(60)
// fmt.Printf("%v\n", c)
// fmt.Printf("%s\n", c)
// }
