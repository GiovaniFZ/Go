package texts

import (
	"fmt"
	"strings"
)

func StringsGo() {
	const str = "Hello, World!"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Contains(str, "Hello"))
}