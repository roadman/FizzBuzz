package main

import "fmt"

func main() {

	for cnt := 1; cnt < 100; cnt++ {
		fizz := cnt % 3
		buzz := cnt % 5
		if fizz == 0 && buzz == 0 {
			fmt.Print("fizzbuzz")
		} else if fizz == 0 {
			fmt.Print("fizz")
		} else if buzz == 0 {
			fmt.Print("buzz")
		} else {
			fmt.Print(cnt)
		}
		fmt.Print(" ")
	}

}
