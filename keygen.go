package keygen // import "ztaylor.me/keygen"

import (
	"math/rand"
	"time"
	"unsafe"
)

// CharSetAlphaNumericCapital is "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var CharSetAlphaNumericCapital = []byte{
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
}

// CharSetAlpha is "abcdefghijklmnopqrstuvwxyz"
var CharSetAlpha = []byte{
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
}

// CharSetCapital is "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var CharSetCapital = []byte{
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
}

// CharSetNumeric is "0123456789"
var CharSetNumeric = []byte{
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
}

// Keygener is an interface type that generates random strings
type Keygener interface {
	Keygen() string
}

// New creates a value quickly
func New(size int, charset []byte, rand *rand.Rand) string {
	key := make([]byte, size)
	for i := 0; i < size; i++ {
		key[i] = charset[int(rand.Int31n(int32(len(charset))))]
	}
	// unsafe trick from strings.Builder.String()
	return *(*string)(unsafe.Pointer(&key))
}

// Settings provides random string generation settings
//
// implements Kengener
type Settings struct {
	KeySize int
	CharSet []byte
	Rand    *rand.Rand
}

// Keygen creates a new Key with the saved Settings
func (settings *Settings) Keygen() string {
	return New(settings.KeySize, settings.CharSet, settings.Rand)
}

// DefaultSettings provides global var for basic functionality
var DefaultSettings = &Settings{
	KeySize: 7,
	CharSet: CharSetAlphaNumericCapital,
	Rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
}

// NewVal creates a key using DefaultSettings
func NewVal() string {
	return DefaultSettings.Keygen()
}
