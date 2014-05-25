package main

import (
	"fmt"
	"github.com/gritty/en"
)

func main() {
	var f float64
	var s string
	var m float64
	var e int

	fmt.Printf("Public \"en\" functions:\n")

	// EntoF() converts an engineering notation number to float64.
	f = en.EntoF(-632.5, en.Nano) // returns 6.325e-07
	fmt.Println("en.EntoF(632.5, en.Nano)    returns:", f)
	// FtoEn() converts a float64 number to its engineering notation.
	s = en.FtoEn(-6.325e-07) // returns "633 n"
	fmt.Println("en.FtoEn(6.325e-07)         returns:", s)
	// FtoME() breaks out a float64 number into its engineering notation
	// mantissa and exponent parts.
	m, e = en.FtoME(-6.325e-07) // returns "633.00", -12
	// m = mantissa, e = exponent
	fmt.Println("en.FtoME(6.325e-07)         returns:", m, e)
	// GetEnCode() returns the engineering notation for a specified
	// exponent.
	s = en.GetEnCode(en.Micro) // returns "µ"
	fmt.Println("en.GetEnCode(en.Micro)      returns:", s)

	fmt.Printf("\nOut of range results:\n")

	// FtoEn() returns a float64 number if it receives a number that is
	// not within an engineering notation range.
	s = en.FtoEn(0.1e-24) // returns
	fmt.Println("en.FtoEn(0.1e-24)           returns:", s)
	s = en.FtoEn(1000e+24) // returns 1.00e+27
	fmt.Println("en.FtoEn(1000e+24)          returns:", s)

	fmt.Printf("\nFour ways to convert 4.83k to a float64:\n")

	f = en.EntoF(4.83, en.Kilo)
	fmt.Printf("en.EntoF(4.83, en.Kilo)     returns: %.3e\n", f)
	fmt.Println("en.FtoEn(4.830e+03)         returns:", en.FtoEn(f))
	f = en.EntoF(0.00483, en.Mega)
	fmt.Printf("en.EntoF(0.00483, en.Mega)  returns: %.3e\n", f)
	fmt.Println("en.FtoEn(4.830e+03)         returns:", en.FtoEn(f))
	f = en.EntoF(4830.0, en.Unit)
	fmt.Printf("en.EntoF(4830.0, en.Unit)   returns: %.3e\n", f)
	fmt.Println("en.FtoEn(4.830e+03)         returns:", en.FtoEn(f))
	f = en.EntoF(4830000, en.Milli)
	fmt.Printf("en.EntoF(4830000, en.Milli) returns: %.3e\n", f)
	fmt.Println("en.FtoEn(4.830e+03)         returns:", en.FtoEn(f))

	fmt.Printf("\n%s\n",
		"Calculated Thevenin voltage and resistance for a circuit:")

	//   +---6kΩ----+---4kΩ--A-+
	//   |          |          |
	//  72V        3kΩ      RloadΩ
	//   |          |          |
	//   +----------+--------B-+

	v1 := float64(72)          // 72V dc
	r1 := en.EntoF(6, en.Kilo) // 6kΩ
	r2 := en.EntoF(3, en.Kilo) // 3kΩ
	r3 := en.EntoF(4, en.Kilo) // 4kΩ

	// calculate Vth
	i := v1 / (r1 + r2)
	fmt.Println("I   =", en.FtoEn(i)+en.Amp)
	Vth := i * r2
	fmt.Println("Vth =", en.FtoEn(Vth)+en.Volt)

	// calculate Rth using product over sum
	Rth := r3 + (r2*r1)/(r2+r1)

	// Thevenin circuit:
	//   +----6kΩ---A-+
	//   |            |
	//  24V         RloadΩ
	//   |            |
	//   +----------B-+

	fmt.Println("Rth =", en.FtoEn(Rth)+en.Ohm)
}
