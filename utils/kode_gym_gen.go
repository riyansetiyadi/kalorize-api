package utils

import "strings"

func GenerateKodeGym(namaGym string) string {
	var kodeGym string
	kodeGym = strings.Split(namaGym, " ")[0]
	return kodeGym + RandomInt(5)
}
