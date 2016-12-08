package main

import (
    "fmt"
    "os"
    "bufio"
    "encoding/json"
    )

type result struct {
   File_name string
   Line    int
   Words  int
   Repeated_words int
}
func main() {

  f, err := os.Open("name.txt")

  if err != nil {

     panic(err)
  }


  scanner := bufio.NewScanner(f)
 
  scanner.Split(bufio.ScanWords)
  
  count := 0
   for scanner.Scan() {
       fmt.Println(scanner.Text())
       count++
   }

   if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading input:", err)
    }
       fmt.Printf("%d\n", count)

 res1D := result{
           File_name:   "name.txt",
           Line     :   10 ,
           Words    :   count ,
           Repeated_words : count }
           res1B, _ := json.Marshal(res1D)
                 fmt.Println(string(res1B))

  /* data := csv.NewReader(f)

   content, err := data.ReadAll()

   if err != nil {
     panic(err)
  }
  var nolines int;
  for _, val := range content {
       fmt.Println( val)
       nolines++;
  }
  fmt.Println("Number of lines are ", nolines) */
} 
