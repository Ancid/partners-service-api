package partner

import (
	"log"
	"math"
	"sort"

	"api_service/database"
)

const radius = 6371 // Earth's mean radius in kilometers

type Service interface {
	GetAllPartners() ([]Partner, error)
	GetSinglePartner(ID int) (PartnerInfo, error)
	GetOrderedPartners(materials []string, latitude float64, longitude float64, limit int) ([]Partner, error)
	getPartnerServices(ID uint) ([]string, error)
}

type service struct {
	repository database.Repository
}

func (s service) GetAllPartners() ([]Partner, error) {
	partners, err := s.repository.GetAllPartners()
	if err != nil {
		return nil, err
	}

	result := make([]Partner, len(partners))
	for i, partner := range partners {
		result[i] = Partner{
			ID:        partner.ID,
			Name:      partner.Name,
			Rating:    partner.Rating,
			Latitude:  partner.Latitude,
			Longitude: partner.Longitude,
			Phone:     partner.Phone,
		}
	}

	return result, nil
}

func (s service) GetSinglePartner(ID int) (PartnerInfo, error) {
	partner, err := s.repository.GetSinglePartner(ID)
	if err != nil {
		return PartnerInfo{}, err
	}

	result := PartnerInfo{
		ID:              partner.ID,
		Name:            partner.Name,
		Rating:          partner.Rating,
		Latitude:        partner.Latitude,
		Longitude:       partner.Longitude,
		Phone:           partner.Phone,
		OperatingRadius: partner.OperatingRadius,
	}

	materials, err := s.getPartnerServices(partner.ID)
	result.Materials = materials

	return result, nil
}

func (s service) getPartnerServices(ID uint) ([]string, error) {
	var result []database.Service

	result, err := s.repository.GetPartnerServices(ID)
	if err != nil {
		return nil, err
	}

	var resultServices []string
	for _, v := range result {
		resultServices = append(resultServices, v.Material)
	}

	return resultServices, nil
}

func (s service) GetOrderedPartners(materials []string, latitude float64, longitude float64, limit int) ([]Partner, error) {
	var result []database.Partner
	var err error
	if len(materials) > 0 {
		result, err = s.repository.GetPartnersByMaterial(materials, limit)
	} else {
		result, err = s.repository.GetAllPartners()
	}

	if err != nil {
		return nil, err
	}

	var partners []database.Partner
	for i, _ := range result {
		distanceToPartner := s.distance(result[i].Latitude, result[i].Longitude, latitude, longitude)
		log.Println("Distance to "+result[i].Name, distanceToPartner)
		//log.Println("Partner", partner)
		log.Println("Operating radius ", result[i].OperatingRadius)
		log.Println("==========")
		if distanceToPartner < result[i].OperatingRadius {
			partners = append(partners, result[i])
		}
	}

	sort.Slice(partners, func(i, j int) bool {
		partnerOne, partnerTwo := partners[i], partners[j]
		distanceToPartnerOne := s.distance(partnerOne.Latitude, partnerOne.Longitude, latitude, longitude)
		distanceToPartnerTwo := s.distance(partnerTwo.Latitude, partnerTwo.Longitude, latitude, longitude)
		switch {
		case partnerOne.Rating != partnerTwo.Rating:
			return partnerOne.Rating < partnerTwo.Rating
		default:
			return distanceToPartnerOne < distanceToPartnerTwo
		}
	})

	resultPartners := make([]Partner, len(partners))
	for i, partner := range partners {
		resultPartners[i] = Partner{
			ID:        partner.ID,
			Name:      partner.Name,
			Rating:    partner.Rating,
			Latitude:  partner.Latitude,
			Longitude: partner.Longitude,
			Phone:     partner.Phone,
		}
	}

	return resultPartners, nil
}

func (s service) calculateDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	degreesLat := degrees2radians(lat2 - lat1)
	degreesLong := degrees2radians(lng2 - lng1)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(lat1))*
			math.Cos(degrees2radians(lat2))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := radius * c

	return d
}

func (s service) distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.853159616

	return dist
}

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func NewService(partnerRepository database.Repository) Service {
	return service{
		repository: partnerRepository,
	}
}
