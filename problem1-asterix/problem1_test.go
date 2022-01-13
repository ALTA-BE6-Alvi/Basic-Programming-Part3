package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayWithAsterix(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, "    * \n   * * \n  * * * \n * * * * \n* * * * * \n", PlayWithAsterix(5), "Hasil harus sama")
	})
}
