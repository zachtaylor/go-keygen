package keygen // import "ztaylor.me/keygen"

import (
	"math/rand"
	"time"
	"unsafe"

	pkg_charset "ztaylor.me/charset"
)

// Keygener is an interface type that generates random strings
type Keygener interface {
	Keygen() string
}

// New creates a value quickly
func New(size int, charset string, rand *rand.Rand) string {
	key := make([]byte, size)
	for i, clen := 0, len(charset); i < size; i++ {
		key[i] = charset[int(rand.Int31n(int32(clen)))]
	}
	// unsafe trick from strings.Builder.String()
	return *(*string)(unsafe.Pointer(&key))
}

// Settings provides random string generation settings
//
// implements Kengener
type Settings struct {
	KeySize int
	CharSet string
	Rand    *rand.Rand
}

// Keygen creates a new Key with the saved Settings
func (settings *Settings) Keygen() string {
	return New(settings.KeySize, settings.CharSet, settings.Rand)
}

// DefaultSettings provides global var for basic functionality
var DefaultSettings = &Settings{
	KeySize: 7,
	CharSet: pkg_charset.AlphaNumericCapital,
	Rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
}

// NewVal creates a key using DefaultSettings
func NewVal() string {
	return DefaultSettings.Keygen()
}
