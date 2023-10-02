package tabungan

import (
	"math/rand"
)

type Service interface {
	CreateTabungan(input CreateTabunganInput) (Tabungan, error)
	CreateMutasi(input CreateMutasiInput) (int, error)
	GetTabunganByNoRek(NoRekening int) (Tabungan, error)
	IsTabunganAvailable(input CreateTabunganInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTabungan(input CreateTabunganInput) (Tabungan, error) {
	min := 10000000
	max := 99999999
	norek := rand.Intn(max-min+1)+min

	tabungan := Tabungan{}
	tabungan.Nama = input.Nama
	tabungan.NIK = input.NIK
	tabungan.NoHP = input.NoHP
	tabungan.NoRekening = norek

	NewTabungan, err := s.repository.Save(tabungan)
	if err != nil {
		return NewTabungan, err
	}

	return NewTabungan, nil
}

func (s *service) IsTabunganAvailable(input CreateTabunganInput) (bool, error) {
	NIK := input.NIK
	NoHP := input.NoHP

	NewTabungan, err := s.repository.FindByNIK(NIK, NoHP)
	if err != nil {
		return false, err
	}

	if NewTabungan.Id == 0 {
		return false,nil
	}

	return true,nil
}

func (s *service) CreateMutasi(input CreateMutasiInput) (int, error) {
	mutasi := Mutasi{}
	mutasi.Nominal = input.Nominal
	mutasi.Type = input.Type

	Tabungan, err := s.repository.FindByNoRek(input.NoRekening)
	if err != nil {
		return 0, err
	}

	mutasi.TabunganID = Tabungan.Id

	NewMutasi, err := s.repository.SaveMutasi(mutasi)
	if err != nil {
		return 0, err
	}

	if (NewMutasi.Type == "C"){
		Tabungan.Saldo = Tabungan.Saldo + input.Nominal
	} else if (NewMutasi.Type == "D") {
		Tabungan.Saldo = Tabungan.Saldo - input.Nominal
	}

	updatedTabungan, err := s.repository.Update(Tabungan)
	if err != nil {
		return 0, err
	}

	return updatedTabungan.Saldo, nil
}


func (s *service) GetTabunganByNoRek(NoRekening int) (Tabungan, error) {
	tabungan, err := s.repository.FindByNoRek(NoRekening)

	if err != nil {
		return tabungan, err
	}

	return tabungan, nil
}