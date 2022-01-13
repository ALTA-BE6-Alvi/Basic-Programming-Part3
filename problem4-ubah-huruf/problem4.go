package main

import "fmt"

func UbahHuruf(sentence string) string {
	// your code here
	var output string

	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 65 && sentence[i] <= 90 {
			number := sentence[i] + 10
			if number < 65 {
				number = 90 - (65 - number)
			} else if number > 90 {
				number = 65 + (number - 90 - 1)
			}
			output += string(number)
		} else if sentence[i] >= 97 && sentence[i] <= 122 {
			number := sentence[i] + 10
			if number < 97 {
				number = 122 - (97 - number)
			} else if number > 122 {
				number = 97 + (number - 122)
			}
			output += string(number)
		} else {
			output += string(sentence[i])
		}
	}

	return output
}

func main() {
	fmt.Println(UbahHuruf("SEPULSA OKE"))     // COZEVCK YUO
	fmt.Println(UbahHuruf("ALTERRA ACADEMY")) // KVDOBBK KMKNOWI
	fmt.Println(UbahHuruf("INDONESIA"))       // SXNYXOCSK
	fmt.Println(UbahHuruf("GOLANG"))          // QYVKXQ
	fmt.Println(UbahHuruf("PROGRAMMER"))      // ZBYQBKWWOB
}
