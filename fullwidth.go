package gullwidth

import (
	"errors"
	"fmt"
	"unicode"
)

// latinASCIICharMapping defines all the latin fullwidth characters that map
// to a character defined in the ASCII set.
//
// Currently Katakana, Hangul, and other special characters are not supported.
//
// U+FF00 does not correspond to a fullwidth ASCII 20 character as this is defined
// by the ideographic space character (defined separately below)
var latinASCIICharMapping map[byte]rune = map[byte]rune{
	'!': '！', '"': '＂', '#': '＃', '$': '＄',
	'%': '％', '&': '＆', '\'': '＇', '(': '（',
	')': '）', '*': '＊', '+': '＋', ',': '，',
	'-': '－', '.': '．', '/': '／', '0': '０',
	'1': '１', '2': '２', '3': '３', '4': '４',
	'5': '５', '6': '６', '7': '７', '8': '８',
	'9': '９', ':': '：', ';': '；', '<': '＜',
	'=': '＝', '>': '＞', '?': '？', '@': '＠',
	'A': 'Ａ', 'B': 'Ｂ', 'C': 'Ｃ', 'D': 'Ｄ',
	'E': 'Ｅ', 'F': 'Ｆ', 'G': 'Ｇ', 'H': 'Ｈ',
	'I': 'Ｉ', 'J': 'Ｊ', 'K': 'Ｋ', 'L': 'Ｌ',
	'M': 'Ｍ', 'N': 'Ｎ', 'O': 'Ｏ', 'P': 'Ｐ',
	'Q': 'Ｑ', 'R': 'Ｒ', 'S': 'Ｓ', 'T': 'Ｔ',
	'U': 'Ｕ', 'V': 'Ｖ', 'W': 'Ｗ', 'X': 'Ｘ',
	'Y': 'Ｙ', 'Z': 'Ｚ', '[': '［', '\\': '＼',
	']': '］', '^': '＾', '_': '＿', '`': '｀',
	'a': 'ａ', 'b': 'ｂ', 'c': 'ｃ', 'd': 'ｄ',
	'e': 'ｅ', 'f': 'ｆ', 'g': 'ｇ', 'h': 'ｈ',
	'i': 'ｉ', 'j': 'ｊ', 'k': 'ｋ', 'l': 'ｌ',
	'm': 'ｍ', 'n': 'ｎ', 'o': 'ｏ', 'p': 'ｐ',
	'q': 'ｑ', 'r': 'ｒ', 's': 'ｓ', 't': 'ｔ',
	'u': 'ｕ', 'v': 'ｖ', 'w': 'ｗ', 'x': 'ｘ',
	'y': 'ｙ', 'z': 'ｚ', '{': '｛', '|': '｜',
	'}': '｝', '~': '～',
}

// ideographicSpaceCharMapping covers fullwidth space, which is provided by
// U+3000 as the "Ideographic Space" character. Defined separately as this
// separate from the fullwidth definitions in unicode.
var ideographicSpaceCharMapping map[byte]rune = map[byte]rune{
	' ': '　',
}

// ErrNotASCII is returned when the Fullwidth function is passed a string
// containing a non-ASCII character.
var ErrNotASCII error = errors.New("string contans non-ascii characters")

// Fullwidth takes an ASCII string and converts it to unicode fullwidth
// characters. If the provided string contains characters that are not
// part of the ASCII character set, an error is returned.
func Fullwidth(ascii string) (fullwidth string, err error) {
	if c, err := isASCII(ascii); err != nil {
		return "", fmt.Errorf("%w: rune is not in the ASCII character set: %c", err, rune(c))
	}

	for i := 0; i < len(ascii); i++ {
		fullwidth = fullwidth + string(latinASCIICharMapping[ascii[i]])
		fullwidth = fullwidth + string(ideographicSpaceCharMapping[ascii[i]])
	}

	return fullwidth, nil
}

// isASCII checks if a given string contains non-ASCII characters. If
// a non-ASCII character is found, it is returned, alongside ErrNotASCII
func isASCII(s string) (byte, error) {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return s[i], ErrNotASCII
		}
	}
	return 0, nil
}
