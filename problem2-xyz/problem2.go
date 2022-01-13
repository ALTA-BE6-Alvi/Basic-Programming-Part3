package main

import "fmt"

func DrawXYZ(n int) string {
	// your code here
	var draw string
	xyz := "XYZ"
	order := 0
	pos := 0

	for i := 0; i < n; i++ {
		for j := 1; j <= n; j++ {
			// Check position
			pos = (i * n) + j
			
			// Check order XYZ
			if pos % 3 == 0 {
				order = 0
			} else if pos % 2 == 0 {
				order = 2
			} else {
				order = 1
			}

			draw += string(xyz[order])
			draw += " "
			
			// constrain order
			if order > 2 {
				order = order - 3
			}
		}
		draw += "\n"
	}

	return draw
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}
