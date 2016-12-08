package main

import (
    "fmt"
    "os"
    "encoding/csv"
    "encoding/json"
    )

func main() {

  f, err := os.Open("name.txt")

  if err != nil {

     panic(err)
  }

  data := csv.NewReader(f)

  content, err := data.ReadAll()

  if err != nil {

     panic(err)
  }
  var nolines int;
  for _, val := range content {
       fmt.Println( val)
       nolines++;
  }
  fmt.Println("Number of lines are ", nolines)
}
