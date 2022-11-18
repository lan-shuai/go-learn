package main

import "fmt"
const (
		B  = iota
		KB = 1 << (10 * iota) // <<移位操作，速度比乘除法快 
		MB = 1 << (10 * iota) // 1<<3 相当于1*2*2*2     0001 -> 1000
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

func main(){
	fmt.Println(B, KB, MB, GB, TB, PB)
}