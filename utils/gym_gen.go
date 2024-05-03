package utils

import (
	"strings"
	"unicode"
)

func GenerateKodeGym(namaGym string) string {
	// make it lowercase
	namaGym = strings.ToLower(namaGym)
	kodeGym := strings.Split(namaGym, " ")[0]
	return kodeGym + RandomInt(5)
}

func GetAlphabetFromCode(code string) string {
	alphabet := ""
	for _, char := range code {
		if unicode.IsLetter(char) {
			alphabet += string(char)
		} else {
			break
		}
	}
	return alphabet
}
