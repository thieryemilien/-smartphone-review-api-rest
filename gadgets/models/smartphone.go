package smartphone

// Smartphone represents a structure for smartphone
type Smartphone struct {
	Id            int64
	Name          string
	Price         float64
	CountryOrigin string
	Os            string
}

// CreateSmartphoneCMD represents a command to create a smartphone
type CreateSmartphoneCMD struct {
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	CountryOrigin string  `json:"country_origin"`
	Os            string  `json:"os"`
}
