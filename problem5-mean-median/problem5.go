package main

import "fmt"

func MeanMedian(arrayInput []float64) (float64, float64) {
	// your code here
	mean := 0.0
	median := 0.0
	sum := 0.0
	arrLenght := float64(len(arrayInput))

	// Mean calculation
	for i := 0; i < len(arrayInput); i++ {
		sum += arrayInput[i] 	
	}

	mean = sum / arrLenght

	var mid_pos int
	// Median Calculation
	if len(arrayInput) % 2 == 0 {
		mid_pos = len(arrayInput) / 2
		num1 := arrayInput[mid_pos - 1]
		num2 := arrayInput[mid_pos]
		median = (num1 + num2) / 2
	} else {
		mid_pos = len(arrayInput) / 2
		median = arrayInput[mid_pos]
	}

	return mean, median
}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))       // 3 3
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))     // 10.4 9
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}
