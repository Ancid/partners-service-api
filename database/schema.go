package database

type Partner struct {
	ID              uint
	Name            string
	Phone           string
	Rating          float64
	Latitude        float64
	Longitude       float64
	OperatingRadius float64
}

type Service struct {
	ID        uint
	PartnerID uint
	Material  string
}
