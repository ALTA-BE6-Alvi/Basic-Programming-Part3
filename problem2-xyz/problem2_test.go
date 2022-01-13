package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawXYZ(t *testing.T) {
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, "Y Z X \nZ Y X \nY Z X \n", DrawXYZ(3), "Hasil parameter 3 harus sama")
	})

	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, "Y Z X Z Y \nX Y Z X Z \nY X Y Z X \nZ Y X Y Z \nX Z Y X Y \n", DrawXYZ(5), "Hasil parameter 5 harus sama")
	})
}
