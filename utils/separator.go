package utils

import (
	"fmt"
	"strings"
)

func ConvertToArrayWithCommaSeparator(String string) []string {
	return strings.Split(String, ". ")
}

func ConvertToArrayWithDotSeparator(String string) []string {
	return strings.Split(String, ".. ")
}

func ConvertToArrayWithDoubleLineSeparator(String string) []string {
	return strings.Split(String, "--")
}

func AddNumbering(input []string) []string {
	var result []string
	for i, item := range input {
		result = append(result, fmt.Sprintf("%d. %s", i+1, item))
	}
	return result
}
