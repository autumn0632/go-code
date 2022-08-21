package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})

	ch1 <- struct{}{}
	for {
		go func() {
			<-ch1
			fmt.Println("1")
			ch2 <- struct{}{}
		}()
		go func() {
			<-ch2
			fmt.Println("2")
			ch3 <- struct{}{}
		}()
		go func() {
			<-ch3
			fmt.Println("3")
			ch4 <- struct{}{}
		}()
		go func() {
			<-ch4
			fmt.Println("4")
			ch1 <- struct{}{}
		}()
		time.Sleep(time.Second)

	}

}
