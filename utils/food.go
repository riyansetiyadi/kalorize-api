package utils

import (
	"math/rand"
	"strconv"
	"time"
)

type MakananRequest struct {
	Nama          string   `json:"nama"`
	Jenis         string   `json:"jenis"`
	Bahan         []string `json:"bahan"`
	CookingStep   []string `json:"cookingStep"`
	Kalori        int      `json:"kalori"`
	ListFranchise []string `json:"listFranchise"`
	Protein       int      `json:"protein"`
}

func GenerateIdMakanan(namaMakanan string) string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(901) + 100
	idMakanan := strconv.Itoa(randomNumber)
	return idMakanan
}
