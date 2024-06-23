package strs

import (
	"testing"
)

var testText = RandText(100000)

func BenchmarkCeasarCyfer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CeasarCyfer("aBcDdEfG", -1284)
	}
}

func BenchmarkCeasarCyferV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CeasarCyfer("aBcDdEfG", -1284)
	}
}
