package main

import "fmt"

func main() {
  Impar := GeneradorNumerosImpares()
	for i := 0; i < 15; i++ {
		fmt.Println(Impar())
	}
}

func GeneradorNumerosImpares()func()int {
	i := int(1)
	return func() int {
		var num = i
		i += 2
		return num
	}
}
