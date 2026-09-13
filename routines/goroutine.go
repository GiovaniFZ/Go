package routines

import (
	"fmt"
	"time"
)

func showMsg() {
	fmt.Println("Hello, I'm from a go routine!")
}

func ExecuteRoutine() {
	fmt.Println("---------------------")
	go showMsg()
	time.Sleep(1 * time.Second) // Wait the routine to run
	fmt.Println("This is executed in the main routine.")
}