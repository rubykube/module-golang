package cipher

import (
	"errors"
	"fmt"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Shift struct {
	key int
}

type Caesar struct{}

type Vigenere struct {
	key string
}

func NewShift(key int) Cipher {
	if key < 1 || key > 25 {
		if key > -1 || key < -25 {
			return nil
		}
	}
	return Shift{key}
}

func NewVigenere(key string) Cipher {
	var status int
	status = 0

	if len(key) < 1 {
		return nil
	}
	for i := 0; i < len(key); i++ {
		if key[i] != 'a' {
			status = 1
		}
		if key[i] > 122 || key[i] < 97 {
			return nil
		}
	}
	if status == 0 {
		return nil
	}
	return Vigenere{key}
}

func NewCaesar() Cipher {
	return Caesar{}
}

func Downcase(str string) (string, error) {
	tmp := ""
	if len(str) < 1 {
		return "error", errors.New("error")
	}
	for i := 0; i < len(str); i++ {
		char := str[i]
		if char < 91 && char > 64 {
			char = char + 32
		}
		if char > 96 && char < 123 {
			tmp += string(char)
		}
	}
	return tmp, nil
}

func (shift Shift) Decode(str string) string {
	cipher := ""
	for i := 0; i < len(str); i++ {
		for str[i] < 97 || str[i] > 122 {
			i++
		}
		char := str[i]
		char = (char - byte(shift.key))
		if char > 122 {
			char = 'a' + (char-1)%122
		}
		if char < 97 {
			char = 'z' - 97%(char+1)
		}
		cipher += string(char)
	}
	return cipher
}

func (shift Shift) Encode(str string) string {
	cipher := ""
	str, err := Downcase(str)
	if err != nil {
		return fmt.Sprint("")
	}
	for i := 0; i < len(str); i++ {
		for str[i] < 97 || str[i] > 122 {
			i++
		}
		char := str[i]
		char = (char + byte(shift.key))
		if char > 122 {
			char = 'a' + (char-1)%122
		}
		if char < 97 {
			char = 'z' - 97%(char+1)
		}
		cipher += string(char)
	}
	return cipher
}

func (v Vigenere) Encode(str string) string {
	leng := len(v.key)
	cipher := ""
	str, err := Downcase(str)
	if err != nil {
		return fmt.Sprint("Error!")
	}
	for i := 0; i < len(str); i++ {
		for str[i] < 97 || str[i] > 122 {
			i++
		}
		char := str[i]
		char = (char + (v.key[i%leng] - 'a'))
		if char > 122 {
			char = 'a' + (char-1)%122
		}
		cipher += string(char)
	}
	return cipher
}

func (v Vigenere) Decode(str string) string {
	leng := len(v.key)
	cipher := ""
	for i := 0; i < len(str); i++ {
		for str[i] < 97 || str[i] > 122 {
			i++
		}
		char := str[i]
		char = (char - (v.key[i%leng] - 'a'))
		if char < 97 {
			char = 'z' - 97%(char+1)
		}
		cipher += string(char)
	}
	return cipher
}

func (c Caesar) Encode(s string) string {
	cipher := ""
	str, err := Downcase(s)
	if err != nil {
		return fmt.Sprint("")
	}
	for i := 0; i < len(str); i++ {
		for str[i] < 97 || str[i] > 122 {
			i++
		}
		char := str[i]
		char += 3
		if char > 122 {
			char = 'a' + (char - 1) - 'z'
		}
		cipher += string(char)
	}
	return cipher
}

func (c Caesar) Decode(str string) string {
	cipher := ""
	for i := 0; i < len(str); i++ {
		for str[i] < 97 || str[i] > 122 {
			i++
		}
		char := str[i]
		char -= 3
		if char < 97 {
			char = 'z' - ('a' - char - 1)
		}
		cipher += string(char)
	}
	return cipher
}
