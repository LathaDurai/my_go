package main
import (
  "fmt"
  "math"
  )
// return value
func plus( a int, b int) int {

  return  a+b
}

// combined int
func  minus( a, b int) int {

  return  a-b
}

// multiple return value
func divide ( a, b int) (int, int) {

//    c:= a/b
//    d:= a%b
  //  return c, d
  return a/b, a%b
}

// variadic 
func sum( a ... int) (int) {
  fmt.Println("sum values are ",a)

  total:=0

  for _, val:= range a {
    total = total+val
  }
  return total
}

//closure
func closure(a int) func() int {
  return func() int {
           a++
           return a
  }
}

//Recursion

func fact(a int) int {

  if a == 0 {
    return 1
  }

  return (a * fact(a-1))
}

func main() {

 ret := plus(1,2)
 fmt.Println("plus is ",ret)

 ret = minus(1, 2)
 fmt.Println("minus is ",ret)

 a, b := divide(6, 2)
 fmt.Println("divide is ",a, b)

 tot := sum(1,2,3,4,5)
 fmt.Println("sum of five num is ", tot)

 incr := closure(1)
 fmt.Println( incr())
 fmt.Println( incr())
 fmt.Println( incr())
 fmt.Println( incr())
 fmt.Println( incr())
 incr1 := closure(1)
 fmt.Println( incr1())
 fmt.Println( incr1())
 fmt.Println( incr1())
 fmt.Println( incr1())
 fmt.Println( incr1())

 fac := fact(7)
 fmt.Println( "Factorial is ", fac)

 ret_func := func( x float64) float64 {
   return math.Sqrt(x)
 }

  fmt.Println("on the go function ", ret_func(2));
  fmt.Println("on the go function ", ret_func(4));
}
