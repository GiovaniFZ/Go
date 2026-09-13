package routines

import (
	"fmt"
	"time"
)

func sayWorld() {
	for i := 0; i < 3; i++ {
		fmt.Println("World")
		time.Sleep(time.Second)
	}
}

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello")
		time.Sleep(time.Second)
	}
}

func Multiple() {
	fmt.Println("---------------------")
	fmt.Println("Multiple Routines")
	// Will not run in order
	go sayHello()
	go sayWorld()
	time.Sleep(10 * time.Second)
}
