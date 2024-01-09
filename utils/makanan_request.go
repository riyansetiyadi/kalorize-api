package utils

type MakananRequest struct {
	Nama        string   `json:"nama"`
	Jenis       string   `json:"jenis"`
	Bahan       []string `json:"bahan"`
	CookingStep []string `json:"cookingStep"`
	Kalori      int      `json:"kalori"`
	Protein     int      `json:"protein"`
}
