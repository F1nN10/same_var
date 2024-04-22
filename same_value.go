package main

import "fmt"

func main() {
	h3ad := `
+-----------------------------------------+
|               Same value?               |
+-----------------------------------------+
	`

	fmt.Println(h3ad)

	data := ""
	fmt.Print("X.Value: ")
	fmt.Scanln(&data)

	div()

	comp := ""
	fmt.Print("Y.Value: ")
	fmt.Scanln(&comp)

	div()

	if data == comp {
		fmt.Println("MATCH!")
	} else {
		fmt.Println("FAIL!")
	}

	fmt.Println("	Bye!")
}

func div() {
	fmt.Println("-----------------------")
}
