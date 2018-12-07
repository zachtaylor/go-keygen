package keygen_test

import (
	"testing"

	"ztaylor.me/keygen"
)

func BenchmarkNewVal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		keygen.NewVal()
	}
}

func BenchmarkNewNumeric3(b *testing.B) {
	rand := NewRand()
	charSet := keygen.CharSetNumeric
	len := 3
	for n := 0; n < b.N; n++ {
		keygen.New(len, charSet, rand)
	}
}

func BenchmarkNewNumeric10(b *testing.B) {
	rand := NewRand()
	charSet := keygen.CharSetNumeric
	len := 10
	for n := 0; n < b.N; n++ {
		keygen.New(len, charSet, rand)
	}
}

func BenchmarkNewNumeric100(b *testing.B) {
	rand := NewRand()
	charSet := keygen.CharSetAlphaNumericCapital
	len := 100
	for n := 0; n < b.N; n++ {
		keygen.New(len, charSet, rand)
	}
}
func BenchmarkNewAlphaNumericCapital3(b *testing.B) {
	rand := NewRand()
	charSet := keygen.CharSetAlphaNumericCapital
	len := 3
	for n := 0; n < b.N; n++ {
		keygen.New(len, charSet, rand)
	}
}

func BenchmarkNewAlphaNumericCapital10(b *testing.B) {
	rand := NewRand()
	charSet := keygen.CharSetAlphaNumericCapital
	len := 10
	for n := 0; n < b.N; n++ {
		keygen.New(len, charSet, rand)
	}
}

func BenchmarkNewAlphaNumericCapital100(b *testing.B) {
	rand := NewRand()
	charSet := keygen.CharSetAlphaNumericCapital
	len := 100
	for n := 0; n < b.N; n++ {
		keygen.New(len, charSet, rand)
	}
}
