package seeds

import (
	"api_service/database"

	"gorm.io/gorm"
)

func CreatePartner(
	db *gorm.DB,
	partnerName string,
	phone string,
	rating float64,
	latitude float64,
	longitude float64,
	radius float64,
	materials []string,
) error {
	partner := database.Partner{
		Name:            partnerName,
		Phone:           phone,
		Rating:          rating,
		Latitude:        latitude,
		Longitude:       longitude,
		OperatingRadius: radius,
	}
	err := db.Create(&partner).Error

	if err != nil {
		return err
	}

	for i, _ := range materials {
		err = db.Create(
			&database.Service{
				PartnerID: partner.ID,
				Material:  materials[i],
			},
		).Error

		if err != nil {
			return err
		}
	}

	return err
}

func All() []Partner {
	return []Partner{
		Partner{
			Name: "Create partner 1",
			Run: func(db *gorm.DB) error {
				return CreatePartner(db, "Dusseldorf partner", "58948008", 10, 51.205378, 6.659214, 500, []string{"wood", "carpet", "tiles"})
			},
		},
		Partner{
			Name: "Create partner 2",
			Run: func(db *gorm.DB) error {
				return CreatePartner(db, "Berlin partner", "84029739", 9, 52.555481, 13.376390, 500, []string{"wood", "tiles"})
			},
		},
		Partner{
			Name: "Create partner 2",
			Run: func(db *gorm.DB) error {
				return CreatePartner(db, "Munchen partner", "239856778", 8, 48.109838, 11.463603, 500, []string{"wood", "tiles"})
			},
		},
	}
}
