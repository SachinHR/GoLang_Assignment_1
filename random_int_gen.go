package main

import (
	"fmt"
	"math/rand"
)

func Randomvalue(output chan int) {
	res := rand.Intn(10)
	output <- res
}

func main() {

	output := make(chan int)
	go Randomvalue(output)
	
	p_out := <-output
	fmt.Println(p_out)
}

