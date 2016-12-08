package main 

import "fmt"

func main () {

   var map1 map[int]string   // cant assign values directly , use make

   fmt.Println("map1 is ", map1, "length is ", len(map1))

   // make assignment to entry in NIL map
   // map1[0] = "latha"
   //  map1[1] = "Baskar"
   
   // so we need to make the map

   map1 = make(map[int]string)
  
   map1[0] = "latha"
   map1[1] = "Baskar"
   fmt.Println("map1 is ", map1, "length is ", len(map1))

   // direct assignment

   map2 := map[string]int {"Latha":1 , "Baskar": 2, "foo":3, "bar": 4}
   fmt.Println("map2 is ", map2, "length is ", len(map2))

   delete(map2, "foo")

   fmt.Println("map2 is ", map2, "length is ", len(map2))

   val ,status := map2["latha"]
   fmt.Println("val is ", val, "status is ", status)
   
   if status { 
     fmt.Println("key is present")
   }   else  {
     fmt.Println("key is not present")
   }
   val, status = map2["Latha"]

   fmt.Println("val is ", val, "status is ", status)

   for key, value := range map1{

     fmt.Println( "key is ",key, "value is ",value)
   }
}
