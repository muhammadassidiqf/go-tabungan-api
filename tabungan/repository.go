package tabungan

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(tabungan Tabungan) (Tabungan, error)
	SaveMutasi(mutasi Mutasi) (Mutasi, error)
	Update(tabungan Tabungan) (Tabungan, error)
	FindByNIK(NIK int, NoHP string) (Tabungan, error)
	FindByNoRek(NoRekening int) (Tabungan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(tabungan Tabungan) (Tabungan, error) {
	err := r.db.Create(&tabungan).Error

	if err != nil {
		return tabungan, err
	}

	return tabungan, nil
}

func (r *repository) Update(tabungan Tabungan) (Tabungan, error) {
	err := r.db.Save(&tabungan).Error

	if err != nil {
		return tabungan, err
	}

	return tabungan, nil
}

func (r *repository) FindByNIK(NIK int, NoHP string) (Tabungan, error) {
	var tabungan Tabungan
	err := r.db.Where("nik = ? or no_hp = ?", NIK, NoHP).First(&tabungan).Error
	if err != nil {
		return tabungan, err
	}

	return tabungan, nil
}

func (r *repository) FindByNoRek(NoRekening int) (Tabungan, error) {
	var tabungan Tabungan
	err := r.db.Where("no_rekening = ?", NoRekening).Preload("Mutasi").Find(&tabungan).Error
	if err != nil {
		return tabungan, err
	}

	return tabungan, nil
}

func (r *repository) SaveMutasi(mutasi Mutasi) (Mutasi, error) {
	err := r.db.Create(&mutasi).Error

	if err != nil {
		return mutasi, err
	}

	return mutasi, nil
}