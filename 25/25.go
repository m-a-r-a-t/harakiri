package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start")

	sleep1(3)

	fmt.Println("sleep1 end")

	sleep2(3)

	fmt.Println("sleep2 end")

}

func sleep1(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func sleep2(x int64) {
	startTime := time.Now().Unix()

	for time.Now().Unix()-startTime < x {
	}
}
