package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	ping := make(chan struct{}, 1)
	pong := make(chan struct{}, 1)
 
	ping<- struct{}{}
 
	go play(ping, pong)
 
	time.Sleep(time.Millisecond)
}
 
func play(ping, pong chan struct{}) {
	for {
		select {
		case <-ping:
			fmt.Println("ping")
			pong<- struct{}{}
		case <-pong:
			fmt.Println("    pong")
			ping<- struct{}{}
		}
	}
}
