package main 
  
import "fmt"
   
type User_Address struct { 
    Name    string 
    city    string 
    Pincode int
} 
  
func main() { 
    a1 := User_Address{"Aksha", "Bangalore", 3623572} 
  
    fmt.Println("User_Address: ", a1) 
}
