package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

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
