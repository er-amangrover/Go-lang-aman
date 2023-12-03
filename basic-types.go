package main

import (
	"fmt"
	"math/cmplx"
)

// bool
// string
// int int8 int16 int32 int 64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alais for uint8

// rune // alais of int32 // unicode represent

// float32 float64

// complex64 complex128

// 00000001 = 1
// 00000010 = 2
// 00000100 = 4
// .
// .
// 10000000 = 2^8 - 1

// 1<<8

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
