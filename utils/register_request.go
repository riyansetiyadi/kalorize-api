package utils

type RegisterRequest struct {
	Fullname             string `json:"fullname"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	JenisKelamin         string `json:"jenis_kelamin"`
	Umur                 int    `json:"umur"`
	BeratBadan           int    `json:"berat_badan"`
	TinggiBadan          int    `json:"tinggi_badan"`
	FrekuensiGym         int    `json:"frekuensi_gym"`
	TargetKalori         int    `json:"target_kalori"`
}
