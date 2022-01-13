package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUbahHuruf(t *testing.T) {
	t.Run("Case: SEPULSA OKE ", func(t *testing.T) {
		assert.Equal(t, "COZEVCK YUO", UbahHuruf("SEPULSA OKE"), "Hasil output tidak sesuai")
	})
	t.Run("Case: ALTERRA ACADEMY ", func(t *testing.T) {
		assert.Equal(t, "KVDOBBK KMKNOWI", UbahHuruf("ALTERRA ACADEMY"), "Hasil output tidak sesuai")
	})
	t.Run("Case: INDONESIA ", func(t *testing.T) {
		assert.Equal(t, "SXNYXOCSK", UbahHuruf("INDONESIA"), "Hasil output tidak sesuai")
	})
	t.Run("Case: GOLANG ", func(t *testing.T) {
		assert.Equal(t, "QYVKXQ", UbahHuruf("GOLANG"), "Hasil output tidak sesuai")
	})
	t.Run("Case: PROGRAMMER ", func(t *testing.T) {
		assert.Equal(t, "ZBYQBKWWOB", UbahHuruf("PROGRAMMER"), "Hasil output tidak sesuai")
	})
}
