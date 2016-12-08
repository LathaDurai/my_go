package main
import "fmt"

func swap(a ,b *int ) {

  var temp int
  temp = *a
  *a = *b
  *b = temp
}

func main() {

  var a int
  a = 10

  var ptr *int
  ptr = &a

  fmt.Println("Value of a is ", a)
  fmt.Println("address of a is ", &a)

  fmt.Println("ptr has ",ptr) 
  fmt.Println("value of ptr  ",*ptr)


  // pointer of array

  var arrptr [3]*int

  aa := 10
  bb:=20
  cc:=30

  arrptr[0] = &aa
  arrptr[1] = &bb
  arrptr[2] = &cc


  fmt.Println("value of arrptr[0]  ",*arrptr[0])
  fmt.Println("value of arrptr[1]  ",*arrptr[1])
  fmt.Println("value of arrptr[2]  ",*arrptr[2])
  *arrptr[0] = 1;
  *arrptr[1] = 2;
  *arrptr[2] = 3;

  fmt.Println("value of arrptr[0]  ",*arrptr[0])
  fmt.Println("value of arrptr[1]  ",*arrptr[1])
  fmt.Println("value of arrptr[2]  ",*arrptr[2])


  //pointer to pointer

  var ptr1 *int
  var ptr2 **int
  var val int = 5

  ptr1 = &val
    ptr2 = &ptr1

  fmt.Println("value of val  ",val)
  fmt.Println("value of ptr1  ",*ptr1)
  fmt.Println("value of ptr2 ",**ptr2)

// pointer to function

  var val1, val2 int
  val1 = 10
  val2 = 20
  fmt.Println("value of val1  ",val1)
  fmt.Println("value of val2 ",val2)
  swap(&val1, &val2)
  fmt.Println("value of val1  ",val1)
  fmt.Println("value of val2 ",val2)
}
