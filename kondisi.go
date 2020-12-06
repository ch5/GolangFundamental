package main

import "fmt"

func main() {
	//var point int = 3
	//
	//if point == 10 {
	//	fmt.Println("lulus dengan nilai sempurna")
	//} else if point > 5 {
	//	fmt.Println("lulus")
	//} else if point == 4 {
	//	fmt.Println("hampir lulus")
	//} else {
	//	fmt.Println("mohon coba ujian kembali")
	//}

	//seleksi kondisi

	var point = 4096.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("Percent %.1f%s Perfect", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("Percent %.1f%s nice", percent, "%")
	} else {
		fmt.Printf("Percent %.1f%s not enough", percent, "%")
	}
}
