package main
import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  "bufio"
  "strings"
//  "bytes"
)

type jsonobject struct {
  E2e E2etype
}

type E2etype struct {
   Global_roles Globalrolestype
   App_roles Approlestype

 }

type Globalrolestype struct {
  Asg_admin Asgadmintype
  Asg_Developer Asgdevelopertype
}

type Asgadmintype struct {
  Users []string
}
type Asgdevelopertype struct {
  Users []string
}

type  Approlestype struct {
  Users []Userstype
}

type Userstype struct {
  Email string
  Roles []string
}

func main() {
  //Read Json file
  file, e := ioutil.ReadFile("asg_iam/e2e_users.json")
  if(e!= nil) {
    fmt.Println("file error", e);
    os.Exit(1)
  }
  //fmt.Println(" output is %s\n", string(file))

  //Place to store decode value 
  var jo jsonobject
  //Decoding
  err := json.Unmarshal(file, &jo)
  if(err != nil) {
    fmt.Println("error is ", err)
  }
  //Printing Values - crosscheck
/*  fmt.Printf("Results: %v\n", jo.E2e.Global_roles.Asg_admin.Users)
  fmt.Printf("Results: %v\n", jo.E2e.Global_roles.Asg_Developer.Users)
  for i := range jo.E2e.App_roles.Users {
    fmt.Printf("Output is %s\n", jo.E2e.App_roles.Users[i].Email)
    for j := range jo.E2e.App_roles.Users[i].Roles {
      fmt.Printf("Output is %s\n", jo.E2e.App_roles.Users[i].Roles[j])
    }
  } */

  //Read Input from User
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter username")
  User_email, _ := reader.ReadString('\n')

  User_email = strings.TrimSpace(User_email)
  fmt.Println("Enter roles")
  User_role, _ := reader.ReadString('\n')
  User_role = strings.TrimSpace(User_role)
  fmt.Println("Input:", User_email, "::", User_role, "#")

  //Check in decoded value
  str1 := "Hello"
  str2 := "Hello"
  if(str1 == str2) {
    fmt.Println("matched")
  } else {
    fmt.Println("ixxx not matched")
  }

  is_email_found := false
  is_role_found := false
  var temp int
 for i := range jo.E2e.App_roles.Users {
    var current_user string = jo.E2e.App_roles.Users[i].Email
    if (strings.EqualFold( User_email, current_user)) { 
      fmt.Println("User Found!!!")
      is_email_found = true
      temp = i
      for j := range jo.E2e.App_roles.Users[i].Roles {
           if (strings.EqualFold( User_role, jo.E2e.App_roles.Users[i].Roles[j])) { 
             fmt.Println("Role Found!!")
             is_role_found = true
           }
        }
      }
    }

    var temp1 int
    var temp2 int
    var temp3 int
    if(is_role_found == false ) && (is_email_found == false ) {
       for i := range jo.E2e.App_roles.Users {
         for j := range jo.E2e.App_roles.Users[i].Roles {
           temp1 = i
           temp2 = j
           fmt.Println("index assigned")
         }
       }
       fmt.Println("Assign new values")
     jo.E2e.App_roles.Users[temp1+1].Email = User_email
     jo.E2e.App_roles.Users[temp1+1].Roles[temp2+1] = User_role
     } else if (is_email_found == true ) && (is_role_found == false ) {
         for j := range jo.E2e.App_roles.Users[temp+1].Roles {
           temp3 = j;
         }
     jo.E2e.App_roles.Users[temp+1].Roles[temp3+1] = User_role
   }
 //   str := `{"Email":User_email, "Roles":[User_role]}`
  //enc := json.NewEncoder(os.Stdout)
  //enc.Encode(str)
 //str := &Userstype{"User_email", "["1", "2"]"}
 // out, err := json.MarshalIndent(str, "", "  ")
 fmt.Printf("Results: %v\n", jo.E2e)
}

