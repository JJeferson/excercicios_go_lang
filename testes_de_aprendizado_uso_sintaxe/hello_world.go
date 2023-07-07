package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("------------------------------")
	/*  %v  é valor  %T é tipo   \n quebra linha*/
	x := 10
	y := "oi"
	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("y: %v, %T \n", y, y)
	fmt.Println("------------------------------")

	for d := 0; d < 10; d++ {
		fmt.Println(d)
	}

	fmt.Println("------------------------------")
	testaFunc("do teste")
	fmt.Println("------------------------------")
	/* se for ter retornoa  função precisa ter um tipo*/
	fmt.Printf("Soma %v", soma(1, 3))

}

func testaFunc(abc string) {
	fmt.Printf("RETORNO %v \n", abc)
}

func soma(a, b int) int {
	return a + b
}
