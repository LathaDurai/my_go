package main

import "fmt"
import "os"
 
func main() {

data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
 log.Fatal(err)
 //panic(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

