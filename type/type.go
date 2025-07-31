package main

import "fmt"

// const value = 10

// func main() {
// 	i := int(value)
// 	f := float64(value)
// 	fmt.Println(i, f)
// }
//

func main() {
	var b byte = 127
	var samallI int32 = 2_147_483_647
	var bigI uint64 = 9_223_372_036_854_775_807
	b += 1
	samallI += 1
	bigI += 1
	fmt.Println(b, samallI, bigI)
}
