package main

import "fmt"

func PlayWithAsterix(n int) string {
	// your code here
	var asterisk string

	for i := 0; i < n; i++ {
		for j := n; j > i + 1; j-- {
			asterisk += " "
		}
		for k := 0; k <= i; k++ {
			asterisk += "* "
		}
		asterisk += "\n"
	}
	return asterisk
}

func main() {
	fmt.Println(PlayWithAsterix(5))
}
