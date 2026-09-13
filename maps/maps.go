package maps

import "fmt"

func CheckPlayers() {
	fmt.Println("---------------------")
	players := map[string]int{
		"giovani": 23,
	}

	value, ok := players["giovani"]
	fmt.Println(value, ok)

	// if value, ok := players["giovani"]; ok {
	// 	fmt.Println(value, ok)
	// }
}
