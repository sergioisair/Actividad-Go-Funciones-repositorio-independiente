package main

import "fmt"

func main() {
  var a int
	var b int
	fmt.Print("Ingrese a: ")
	fmt.Scan(&a)
	fmt.Print("Ingrese b: ")
	fmt.Scan(&b)
	intercambia(&a, &b)
	fmt.Println("Valor a =", a, "\nValor b =", b)
}

func intercambia(a *int, b *int) {
	aux := *b
	*b = *a
	*a = aux
}
