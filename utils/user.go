package utils

type UserRequest struct {
	IdUser               string `json:"id_user"`
	Fullname             string `json:"fullname"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	JenisKelamin         int    `json:"jenis_kelamin"`
	Role                 string `json:"role"`
	ReferalCode          string `json:"referal_code"`
	Umur                 int    `json:"umur"`
	BeratBadan           int    `json:"berat_badan"`
	TinggiBadan          int    `json:"tinggi_badan"`
	FrekuensiGym         int    `json:"frekuensi_gym"`
	TargetKalori         int    `json:"target_kalori"`
	Foto                 string `json:"foto"`
}
