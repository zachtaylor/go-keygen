package keygen // import "ztaylor.me/keygen"

import (
	"math/rand"
	"strings"
	"time"
)

var CharSetAlphaNumericCapital = []byte{
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
}

const KeySize = 7

func NewVal() string {
	return New(KeySize, CharSetAlphaNumericCapital)
}

var New = func(size int, charset []byte) string {
	v := strings.Builder{}
	v.Grow(size)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ringp := 0
	for i := 0; i < size; i++ {
		ringp = (ringp + rand.Int()) % len(charset)
		v.WriteByte(charset[ringp])
	}
	return v.String()
}
