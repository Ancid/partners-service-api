package database

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllPartners() ([]Partner, error)
	GetSinglePartner(ID int) (Partner, error)
	GetPartnersByMaterial(materials []string, limit int) ([]Partner, error)
	GetPartnerServices(ID uint) ([]Service, error)
}

type repository struct {
	db *gorm.DB
}

func (t repository) GetAllPartners() ([]Partner, error) {
	var result []Partner
	err := t.db.Model(&Partner{}).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t repository) GetSinglePartner(ID int) (Partner, error) {
	var result Partner
	err := t.db.Model(&Partner{}).First(&result, ID).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
func (t repository) GetPartnerServices(ID uint) ([]Service, error) {
	var result []Service
	err := t.db.Model(&Service{}).Where("partner_id = ?", ID).Find(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

func (t repository) GetPartnersByMaterial(materials []string, limit int) ([]Partner, error) {
	var partners []Partner
	if err := t.db.Model(&Partner{}).Select("partners.*").Distinct().Limit(limit).
		Joins("inner join services on partners.id = services.partner_id AND services.material IN ?", materials).
		Scan(&partners).Error; err != nil {
		return nil, err
	}

	return partners, nil
}

func NewRepository(db *gorm.DB) Repository {
	return repository{
		db: db,
	}
}
