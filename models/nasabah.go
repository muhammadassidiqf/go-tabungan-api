package models

type Nasabah struct {
    Id     int    `json:"id" gorm:"primaryKey"`
    Nama   string `json:"nama"`
    NIK    int64  `json:"nik"`
    NoHP   int64  `json:"no_hp"`
}