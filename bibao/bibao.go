package main

import "fmt"

func adder() func(int) int {
	sum := 0 //自由变量
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()

	//运行一次，sum保存了值5， 下次会一直带着这次更新的sum值5
/*	i := 5
	fmt.Printf("0+1....+ %d= %d\n", i, a(i))*/

	for i := 0; i < 10; i++ {
		fmt.Printf("0+1....+ %d= %d\n", i, a(i))
	}

	i := 5
	fmt.Printf("0+1....+ %d= %d\n", i, a(i))
}
