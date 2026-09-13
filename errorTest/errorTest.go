package errorTest

import (
	"errors"
	"fmt"
)

func ErrorTest() {
	fmt.Println("---------------------")
	if err := throwAnError(); err != nil {
		fmt.Println(err.Error())
	}
}

func throwAnError() error {
	return errors.New("This is an error")
}
