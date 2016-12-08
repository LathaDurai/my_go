package main
import "fmt"

func main() {
/*********For loop**********/
i :=1
     for i<=4 {
       fmt.Println(i)
         i++
     }

   for i=1; i<5; i++ {

       fmt.Println(i)
   }

   for {
       fmt.Println("end")
       break
   }

   if num :=9 ; num <10 {
     fmt.Println(num)
   }
  
   /************array***********/
   var z [2] float32
   var a =  [5]int {1,2,3,4,5}

   b:= [6]int{11, 22, 33, 44, 55, 66}
  
   fmt.Println("z has : ", z)
   fmt.Println("a has : ", a)
   fmt.Println("b has : ", b)
   
   fmt.Println("a has len: ", len(a) ,"b has len: ", len(b))

   // 2d

   var twoD [2][3]int
 
   fmt.Println("twoD has : ", twoD)

   for i:=0; i< 2 ; i++  {
      for j := 0 ; j <3; j++ {
          twoD[i][j] = j
      }
   }

   fmt.Println("twoD has : ", twoD)


}
