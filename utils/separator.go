package utils

import "strings"

func ConvertToArrayWithCommaSeparator(String string) []string {
	return strings.Split(String, ", ")
}

func ConvertToArrayWithDotSeparator(String string) []string {
	return strings.Split(String, "., ")
}
