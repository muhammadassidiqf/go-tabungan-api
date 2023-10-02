package tabungan


type CreateTabunganInput struct {
	Nama	string `json:"nama" binding:"required"`
	NIK 	int `json:"nik" binding:"required"`
	NoHP  	string `json:"no_hp" binding:"required"`
}

type CreateMutasiInput struct {
	NoRekening	int `json:"no_rekening" binding:"required"`
	Nominal 	int `json:"nominal" binding:"required"`
	Type 		string `json:"type"`
	TabunganID 	int `json:"tabungan_id"`
}

type GetTabunganDetailInput struct {
	NoRekening int `uri:"no_rekening" binding:"required"`
}