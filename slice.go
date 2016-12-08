package main
import "fmt"

func main() {
   /***************slice***********/

   //define array without slice
   // can define using make with len & cap or len alone

   var a1[] int
   var a2 = make([]string, 2);
   a3 := make([]string, 3,4);

   var a4 = []int{1,2,3,4,5,6}

   fmt.Println("a1: " ,a1, "len: ", len(a1), "cap(a1): ",cap(a1))
   fmt.Println("a2: " ,a2, "len: ", len(a2), "cap(a2): ",cap(a2))
   fmt.Println("a3: " ,a3, "len: ", len(a3), "cap(a3): ",cap(a3))
   fmt.Println("a4: " ,a4, "len: ", len(a4), "cap(a4): ",cap(a4))
   
   a2[0] = "latha"
   a2[1] = "Baskar"

   a3[0] = "ll"
   a3[1] = "aa"
   
   fmt.Println("a2: " ,a2, "len: ", len(a2), "cap(a2): ",cap(a2))
   fmt.Println("a3: " ,a3, "len: ", len(a3), "cap(a3): ",cap(a3))
   // nill

   if a1 == nil {
      fmt.Println(" a1 is NILL")
   }

   // bounds
   fmt.Println("a4[2:4] " ,a4[2:4])
   fmt.Println("a4[2:] " ,a4[2:])
   fmt.Println("a4[:4] " ,a4[:4])

   bound := a4[3:5]
   fmt.Println("bound has  " ,bound)


   // append - alwys has return val

   a1 = append(a1, 0)
   a1 = append(a1, 1)

   // copy

   copy(a2, a3)

   fmt.Println("a2: " ,a2, "len: ", len(a2), "cap(a2): ",cap(a2))


   // 2d 
   // length of inner array is varry unlikearrays
   var two [][] int

   fmt.Println("two: " ,two, "len: ", len(two), "cap(two): ",cap(two))

   twoD := make([][]int, 3, 3)


       k:= 5  
     for i:= 0; i< 3; i++{
       twoD[i] = make([]int, k)
       for j:= 0; j<k ; j++ {

          twoD[i][j] = i+j
       }
       k++
     }

   fmt.Println("twoD: " ,twoD, "len: ", len(twoD), "cap(twoD): ",cap(twoD))

   // range

   for index, val := range a2 {
     fmt.Println( index, val)
   }
}
