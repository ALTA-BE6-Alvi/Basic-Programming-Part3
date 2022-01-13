package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeanMedian(t *testing.T) {
	t.Run("Case: [1, 2,3, 4]", func(t *testing.T) {
		mean, median := MeanMedian([]float64{1, 2, 3, 4})
		assert.Equal(t, 2.5, mean, "Hasil output MEAN tidak sesuai")
		assert.Equal(t, 2.5, median, "Hasil output MEAN tidak sesuai")
	})
	t.Run("Case: [1, 2, 3, 4, 5]", func(t *testing.T) {
		mean, median := MeanMedian([]float64{1, 2, 3, 4, 5})
		assert.Equal(t, float64(3), mean, "Hasil output MEAN tidak sesuai")
		assert.Equal(t, float64(3), median, "Hasil output MEAN tidak sesuai")
	})
	t.Run("Case: [7, 8, 9, 13, 15]", func(t *testing.T) {
		mean, median := MeanMedian([]float64{7, 8, 9, 13, 15})
		assert.Equal(t, 10.4, mean, "Hasil output MEAN tidak sesuai")
		assert.Equal(t, float64(9), median, "Hasil output MEAN tidak sesuai")
	})
	t.Run("Case: [10, 20, 30, 40, 50]", func(t *testing.T) {
		mean, median := MeanMedian([]float64{10, 20, 30, 40, 50})
		assert.Equal(t, float64(30), mean, "Hasil output MEAN tidak sesuai")
		assert.Equal(t, float64(30), median, "Hasil output MEAN tidak sesuai")
	})
	t.Run("Case: [15, 20, 30, 60, 120]", func(t *testing.T) {
		mean, median := MeanMedian([]float64{15, 20, 30, 60, 120})
		assert.Equal(t, float64(49), mean, "Hasil output MEAN tidak sesuai")
		assert.Equal(t, float64(30), median, "Hasil output MEAN tidak sesuai")
	})
}
