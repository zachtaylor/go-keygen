package keygen_test

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"ztaylor.me/keygen"
)

// NewRand creates *rand.Rand with source time.Now
func NewRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Tests

func TestRandomnessContainsFirstChar(t *testing.T) {
	rand := NewRand()
	var ok bool
	for i := 0; i < 7; i++ {
		key := keygen.New(64, keygen.CharSetNumeric, rand)
		if strings.Contains(key, "0") {
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
		key := keygen.New(64, keygen.CharSetAlphaNumericCapital, rand)
		if strings.Contains(key, "Z") {
			ok = true
		}
	}
	if !ok {
		t.Fail()
	}
}
