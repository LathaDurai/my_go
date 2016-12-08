
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Enter Input \n")
      var user string
      for user != "q" {
        scanner.Scan()
        user = scanner.Text()
        if(user != "q")  {
          fmt.Println( "Output is ", user)
        }
      } 
