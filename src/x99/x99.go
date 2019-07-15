package main

import "fmt"

func main() {
	//99乘法表
	/**

	//1
	*
	**
	***
	****

	//2
	   *
	  **
	 ***
	****

	//3
	****
	 ***
	  **
	   *

	//4
	****
	***
	**
	*

	 */

	//1
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
			//fmt.Printf("%dx%d=%2d ", i, j, i*j)
		}
		fmt.Printf("\n")
	}
	fmt.Println()

	//2
	for i := 1; i <= 9; i++ {
		for k := 8; k >= i; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Printf("\n")
	}
	fmt.Println()

	for i := 1; i <= 9; i++ {
		for k := 1; k <= i-1; k++ {
			fmt.Print(" ")
		}
		for j := 10-i; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 1; i <= 9; i++ {
		for j := 9; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
