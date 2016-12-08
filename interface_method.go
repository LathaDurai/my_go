package main

import (
    "fmt"
    "math"
    )

type circle struct {
   rad float64
}

type rectangle struct {
   width, height int
}

func (rec rectangle) area() int {

  return rec.width *rec.height
}

func (cir circle) area() float64 {

  return math.Pi * cir.rad * cir.rad
}

func (rec rectangle) perim() int {

  return rec.width +rec.height
}

func (cir circle) perim() float64 {

  return math.Pi * cir.rad * cir.rad
}
func main() {

  r := rectangle{width:3, height:4}
  c := circle{rad: 3}
  res := r.area()
  fmt.Println("area of rec is ", res)
  res = c.area()
  fmt.Println("area of circle is ", res)
  res = r.perim()
  fmt.Println("perim of rec is ", res)
  res = c.perim()
  fmt.Println("perim of circle is ", res)
}
