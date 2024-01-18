package formatter

import (
	"kalorize-api/domain/models"
	"kalorize-api/utils"
)

type MakananFormat struct {
	ID          string
	Nama        string
	Jenis       string
	Bahan       []string
	CookingStep []string
	Kalori      int
	Protein     int
	Foto        string
}

func FormatterMakananLuarIndo(makanan models.Makanan) MakananFormat {
	var makananFormatted MakananFormat
	makananFormatted.ID = makanan.IdMakanan
	makananFormatted.Nama = makanan.Nama
	makananFormatted.Jenis = makanan.Jenis
	makananFormatted.Bahan = utils.ConvertToArrayWithCommaSeparator(makanan.Bahan)
	makananFormatted.Bahan = utils.AddNumbering(makananFormatted.Bahan)
	makananFormatted.CookingStep = utils.ConvertToArrayWithDotSeparator(makanan.CookingStep)
	makananFormatted.CookingStep = utils.AddNumbering(makananFormatted.CookingStep)
	makananFormatted.Kalori = makanan.Kalori
	makananFormatted.Protein = makanan.Protein
	makananFormatted.Foto = makanan.Foto
	return makananFormatted
}

func FormatterMakananIndo(makanan models.Makanan) MakananFormat {
	var makananFormatted MakananFormat
	makananFormatted.ID = makanan.IdMakanan
	makananFormatted.Nama = makanan.Nama
	makananFormatted.Jenis = makanan.Jenis
	makananFormatted.Bahan = utils.ConvertToArrayWithDoubleLineSeparator(makanan.Bahan)
	makananFormatted.Bahan = utils.AddNumbering(makananFormatted.Bahan)
	makananFormatted.CookingStep = utils.ConvertToArrayWithDoubleLineSeparator(makanan.CookingStep)
	makananFormatted.CookingStep = utils.AddNumbering(makananFormatted.CookingStep)
	makananFormatted.Kalori = makanan.Kalori
	makananFormatted.Protein = makanan.Protein
	makananFormatted.Foto = makanan.Foto
	return makananFormatted
}
