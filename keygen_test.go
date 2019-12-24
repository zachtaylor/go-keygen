package keygen_test

import (
	"math/rand"
	"testing"

	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/keygen"
)

// NewRand creates *rand.Rand with source time.Now
func NewRand() *rand.Rand {
	return rand.New(rand.NewSource(cast.Now().UnixNano()))
}

// Tests

func TestRandomnessContainsFirstChar(t *testing.T) {
	rand := NewRand()
	var ok bool
	for i := 0; i < 7; i++ {
		key := keygen.New(64, charset.Numeric, rand)
		if cast.Contains(key, "0") {
			ok = true
		}
	}
	if !ok {
		t.Fail()
	}
}
func TestRandomnessContainsLastChar(t *testing.T) {
	rand := NewRand()
	var ok bool
	for i := 0; i < 7; i++ {
		key := keygen.New(64, charset.Capital, rand)
		if cast.Contains(key, "Z") {
			ok = true
		}
	}
	if !ok {
		t.Fail()
	}
}
