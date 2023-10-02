package tabungan

import "time"

type Tabungan struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Nama       string `json:"nama"`
	NIK        int    `json:"nik"`
	Saldo      int    `json:"saldo"`
	NoHP       string `json:"no_hp"`
	NoRekening int    `json:"no_rekening"`
	Mutasi     []Mutasi
	CreatedAt  time.Time 
	UpdatedAt  time.Time
}

type Mutasi struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	TabunganID int       `json:"tabungan_id"`
	Type       string    `json:"type"`
	Nominal    int       `json:"nominal"`
	CreatedAt  time.Time 
	UpdatedAt  time.Time
}