package schema

type Service struct {
	Thing

	ServiceType string          `json:"serviceType,omitempty"`
	Provider    *Organisation   `json:"provider,omitempty"`
	AreaServed  []string        `json:"areaServed,omitempty"`
	Offers      []Offer         `json:"offers,omitempty"`
}

type ProfessionalService struct {
	Thing

	Image       string   `json:"image,omitempty"`
	Telephone   string   `json:"telephone,omitempty"`
	PriceRange  string   `json:"priceRange,omitempty"`
	AreaServed  []string `json:"areaServed,omitempty"`
	SameAs      []string `json:"sameAs,omitempty"`
}

type Offer struct {
	Node

	URL          string `json:"url,omitempty"`
	Price        string `json:"price,omitempty"`
	PriceCurrency string `json:"priceCurrency,omitempty"`
}

func NewService(name string, options ...NodeOption) Service {
	return Service{
		Thing: Thing{
			Node: newNode(TypeService, options...),
			Name: name,
		},
	}
}

func NewProfessionalService(name string, options ...NodeOption) ProfessionalService {
	return ProfessionalService{
		Thing: Thing{
			Node: newNode(TypeProfessionalService, options...),
			Name: name,
		},
	}
}

func NewOffer(options ...NodeOption) Offer {
	return Offer{
		Node: newNode(TypeOffer, options...),
	}
}
