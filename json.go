package main

import (
 "fmt"
 "encoding/json"
 "os"
 "io/ioutil"
 "time"
 )


type jsonobject struct {
  Name  string
  Environments []string
  SelfLink string
  LastUpdatedDate time.Time
}


func main() {
  file, e := ioutil.ReadFile("env.json")

  if(e != nil) {
    fmt.Println("file error", e)
    os.Exit(1)
  }
  fmt.Printf("%s\n", file) 

  var  jo []jsonobject

  err := json.Unmarshal(file, &jo)
  if(err != nil) {
    fmt.Println("error is ", err)
  }

  for i := range jo {
    fmt.Printf("name is %s\n", jo[i].Name)
    fmt.Printf("Env is  %v\n", jo[i].Environments[0])
    fmt.Printf("Env is  %v\n", jo[i].Environments[1])
    fmt.Printf("Env is  %v\n", jo[i].Environments[2])
    fmt.Printf("Env is  %v\n", jo[i].Environments[3])
    fmt.Printf("SelfLink is  %v\n", jo[i].SelfLink)
    fmt.Printf("LastUpdatedTime %v\n", jo[i].LastUpdatedDate)
  }
  for i := range jo {
    name := jo[i].Name
    for j := range jo[i].Environments {
      fmt.Printf(" %s-%v, ", name,jo[i].Environments[j])
    }
      fmt.Printf("\n")
  }
}
