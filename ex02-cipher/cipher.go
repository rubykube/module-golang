package cipher

import "strings"

type shift int
type vig string

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(sh int) Cipher {
	val := shift(sh)
	switch {
	case val > -26 && val < 0:
		val += 26
		return val
	case val > 0 && val < 26:
		return val
	}
	return nil
}

func encode(value rune, sh int) rune {
	if value >= 97 && value <= 122 {
		return (value-97+rune(sh))%26 + 97
	}

	return -1
}

func decode(value rune, sh int) rune {
	if value >= 97 && value <= 122 {
		return (value-97+rune(26-sh))%26 + 97
	}

	return -1
}

func Downcase(line string) string {
	s := make([]byte, len(line))

	for i := range s {
		symbol := line[i]

		if symbol >= 65 && symbol <= 90 {
			symbol += 32
		}

		s[i] = symbol
	}

	return string(s)
}

func (sh shift) Encode(str string) string {

	lowercase := Downcase(str)

	return strings.Map(
		func(r rune) rune { return encode(r, int(sh)) }, lowercase)
}

func (sh shift) Decode(str string) string {
	lowercase := Downcase(str)
	return strings.Map(
		func(r rune) rune { return decode(r, int(sh)) }, lowercase)
}

func NewVigenere(key string) Cipher {
	flag := false

	for _, value := range key {
		if value < 97 || value > 122 {
			return nil
		}
		if value > 97 {
			flag = true
		}
	}
	if !flag {
		return nil
	}

	return vig(key)
}
func (vign vig) Encode(str string) string {
	lowercase := Downcase(str)
	k := 0
	return strings.Map(
		func(r rune) rune {
			if r = encode(r, int(vign[k]-97)); r >= 0 {
				k = (k + 1) % len(vign)
			}
			return r
		}, lowercase)
}

func (vign vig) Decode(str string) string {
	lowercase := Downcase(str)
	k := 0
	return strings.Map(
		func(r rune) rune {
			if r = decode(r, int(vign[k]-97)); r >= 0 {
				k = (k + 1) % len(vign)
			}
			return r
		}, lowercase)
}
