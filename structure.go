package main

import "fmt"

type person struct {

  name string
  age int
}

func main() {

  //var p1 person = {"latha", 25}

  fmt.Println(person {"baskar", 27})

  var p1 person
  p1.name = "foo"
  p1.age = 11111
  fmt.Println("Name: ", p1.name, "Age: ", p1.age)

  var sptr *person
  sptr = &p1
  fmt.Println("Name: ", sptr.name, "Age: ", sptr.age)

  test := &p1
  fmt.Println("Name: ", test.name, "Age: ", test.age)

}
