package formatter

import (
	"kalorize-api/domain/models"
	"strconv"
)

func FormatterMakananToMultiDimentionalArray(makanan []models.Makanan) [][]string {
	var header = []string{"Nama", "Jenis", "Bahan", "Cooking Step", "Kalori", "Protein"}

	var result [][]string
	result = append(result, header)

	for i := range makanan {
		var row []string
		row = append(row, makanan[i].Nama)
		row = append(row, makanan[i].Jenis)
		row = append(row, makanan[i].Bahan)
		row = append(row, makanan[i].CookingStep)
		row = append(row, intToString(makanan[i].Kalori))
		row = append(row, intToString(makanan[i].Protein))

		result = append(result, row)
	}
	return result
}

func intToString(num int) string {
	return strconv.Itoa(num)
}
