package main

import "fmt"

func main() {
    var s []int
    var n int
    var x int
    fmt.Print("Ingrese la cantidad de números a ingresar: ")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
      fmt.Scan(&x)
      s = append(s, x)
    }
    sum(s...)
}

func sum(nums ...int) {
    var mayor int
    for j := 0; j < len(nums) ; j++ {
        if mayor < nums[j]{
          mayor = nums[j]
        }
    }
    fmt.Println("El mayor es" , mayor)
}
