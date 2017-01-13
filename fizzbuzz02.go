package main

import "fmt"

func main() {
	var fizz int
	var buzz int
	for no := 1; no <= 100; no++ {
		fizz++
		buzz++
		if fizz == 3 && buzz == 5 {
			fmt.Print("fizzbuzz ")
			fizz = 0
			buzz = 0
		} else if fizz == 3 {
			fmt.Print("fizz ")
			fizz = 0
		} else if buzz == 5 {
			fmt.Print("buzz ")
			buzz = 0
		} else {
			fmt.Printf("%d ", no)
		}
	}
}
