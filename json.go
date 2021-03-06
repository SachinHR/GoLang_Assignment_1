package main 
   
import ( 
    "fmt"
    "encoding/json"
) 
    
type Human struct{ 
    Name string 
    Age int
    Address string 
} 
    
func main() {
    man := Human{"Ankit", 23, "Delhi"}
    man_enc, err := json.Marshal(man) 
       
    if err != nil { 
        fmt.Println(err) 
    } 
      
    fmt.Println(string(man_enc)) 
}
