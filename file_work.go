package main

import (
    "fmt"
    "io/ioutil"
)

func main() {

    mydata := []byte("My file contents")

    err := ioutil.WriteFile("myfile.txt", mydata, 0777)

    if err != nil {
        fmt.Println(err)
    }

    data, err := ioutil.ReadFile("myfile.txt")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

}
