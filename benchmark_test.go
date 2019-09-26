package keygen_test

import (
	"testing"

	"ztaylor.me/charset"
	"ztaylor.me/keygen"
)

func BenchmarkNewVal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		keygen.NewVal()
	}
}

func BenchmarkNewNumeric3(b *testing.B) {
	rand := NewRand()
	len := 3
	for n := 0; n < b.N; n++ {
		keygen.New(len, charset.Numeric, rand)
	}
}

func BenchmarkNewNumeric10(b *testing.B) {
	rand := NewRand()
	len := 10
	for n := 0; n < b.N; n++ {
		keygen.New(len, charset.Numeric, rand)
	}
}

func BenchmarkNewNumeric100(b *testing.B) {
	rand := NewRand()
	len := 100
	for n := 0; n < b.N; n++ {
		keygen.New(len, charset.Numeric, rand)
	}
}
func BenchmarkNewAlphaNumericCapital3(b *testing.B) {
	rand := NewRand()
	len := 3
	for n := 0; n < b.N; n++ {
		keygen.New(len, charset.AlphaNumericCapital, rand)
	}
}

func BenchmarkNewAlphaNumericCapital10(b *testing.B) {
	rand := NewRand()
	len := 10
	for n := 0; n < b.N; n++ {
		keygen.New(len, charset.AlphaNumericCapital, rand)
	}
}

func BenchmarkNewAlphaNumericCapital100(b *testing.B) {
	rand := NewRand()
	len := 100
	for n := 0; n < b.N; n++ {
		keygen.New(len, charset.AlphaNumericCapital, rand)
	}
}
