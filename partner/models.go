package partner

type Partner struct {
	ID        uint
	Name      string
	Phone     string
	Rating    float64
	Longitude float64
	Latitude  float64
}

type PartnerInfo struct {
	ID              uint
	Name            string
	Phone           string
	Rating          float64
	Longitude       float64
	Latitude        float64
	OperatingRadius float64
	Materials       []string
}

type PartnerService struct {
	ID       uint
	Material string
}
