package keygen // import "ztaylor.me/keygen"

import (
	"math/rand"

	"ztaylor.me/cast"
	pkg_charset "ztaylor.me/cast/charset"
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
	return cast.StringBytes(key)
}

// NewAlpha uses New with charset.Alpha and DefaultSettings.Rand
func NewAlpha(size int) string {
	return New(size, pkg_charset.Alpha, DefaultSettings.Rand)
}

// NewAlphaCapital uses New with charset.AlphaCapital and DefaultSettings.Rand
func NewAlphaCapital(size int) string {
	return New(size, pkg_charset.AlphaCapital, DefaultSettings.Rand)
}

// NewAlphaCapitalNumeric uses New with charset.AlphaCapitalNumeric and DefaultSettings.Rand
func NewAlphaCapitalNumeric(size int) string {
	return New(size, pkg_charset.AlphaCapitalNumeric, DefaultSettings.Rand)
}

// NewCapital uses New with charset.Capital and DefaultSettings.Rand
func NewCapital(size int) string {
	return New(size, pkg_charset.Capital, DefaultSettings.Rand)
}

// NewCapitalNumeric uses New with charset.CapitalNumeric and DefaultSettings.Rand
func NewCapitalNumeric(size int) string {
	return New(size, pkg_charset.CapitalNumeric, DefaultSettings.Rand)
}

// NewNumeric uses New with charset.Numeric and DefaultSettings.Rand
func NewNumeric(size int) string {
	return New(size, pkg_charset.Numeric, DefaultSettings.Rand)

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
	CharSet: pkg_charset.AlphaCapitalNumeric,
	Rand:    rand.New(rand.NewSource(cast.Now().UnixNano())),
}

// NewVal creates a key using DefaultSettings
func NewVal() string {
	return DefaultSettings.Keygen()
}
