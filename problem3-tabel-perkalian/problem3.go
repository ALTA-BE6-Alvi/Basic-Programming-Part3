package main

import "fmt"
import "strconv"

func CetakTabelPerkalian(number int) string {
	// your code here
	var hasil string
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			hasil += strconv.Itoa(i * j)
			hasil += "\t"
		}
		hasil += "\n"
	}
	
	return hasil
}

func main() {
	fmt.Println(CetakTabelPerkalian(9))
}
