package main

import "fmt"

func main() {
	//var point = 8
	//
	//switch point {
	//case 8:
	//	fmt.Println("Perfect nuber")
	//
	//case 4, 5, 6, 7:
	//	fmt.Println("Nice")
	//
	//default:
	//	fmt.Println("Not Bad")
	//	fmt.Println("You can be better")
	//}

	var point = 2
	switch {
	case point == 8:
		fmt.Println("Perfect score")

	case (point < 8) && (point > 3):
		fmt.Println("Awesome number")

	default:
		fmt.Println("not bad")
		fmt.Println("enough to learn more")
	}
}
