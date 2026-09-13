package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	const f float32 = 42.33
	var i int = 1
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println(needFloat(Big))
}
