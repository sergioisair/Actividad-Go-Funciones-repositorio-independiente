package main

import "fmt"

func main() {
   var num uint
   fmt.Print("Ingrese el número de secuencias fibonacci que desee calcular: ")
   fmt.Scan(&num)
   fmt.Println("El resultado es ", fibonacci(num))
}

func fibonacci(i uint) uint {
   if i <= 1 {
      return i
   }
   return fibonacci(i-1) + fibonacci(i-2)
}
