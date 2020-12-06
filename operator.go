package main

import "fmt"

func main() {
	var left = false
	var right = true

	var lefAndRight = left && right
	fmt.Printf("left && right \t(%t)\n", lefAndRight)

	var leftorRight = left || right
	fmt.Printf("left || right \t(%t)\n", leftorRight)

	var leftReverse = !left
	fmt.Printf("!left \t(%t)\n", leftReverse)

}
