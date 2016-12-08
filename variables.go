package main

import "fmt"
import "math"
const ct = 10000;
func main() {

  var lat string = "lathaduraisamy"
  fmt.Println(lat)
  
  var a, b, c int = 1,2,3
  fmt.Println(a, b, c)
  
  var bb bool = false
  fmt.Println(bb)
  
  d := "autodeclarate"
  fmt.Println(d)

  var e float32
  fmt.Println(e)

//Const

  fmt.Println(ct)
  
  const aq =10.1

  fmt.Println(aq)

  fmt.Println(int64 (aq))   //explicit type conversion

  //aq = 11.10   //error

  fmt.Println( math.Sin(aq))

}

