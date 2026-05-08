package demo

type CustomerSpec struct {
	ID        string
	Name      string
	SMSNumber string
}

type StoreSpec struct {
	ID            string
	Name          string
	Location      string
	Participating bool
	Products      []ProductSpec
}

type ProductSpec struct {
	ID          string
	Name        string
	Description string
	SKU         string
	Price       float64
}

type BootstrapSpec struct {
	Customer CustomerSpec
	Stores   []StoreSpec
}

func Spec() BootstrapSpec {
	return BootstrapSpec{
		Customer: CustomerSpec{
			ID:        "11111111-1111-1111-1111-111111111111",
			Name:      "Demo Shopper",
			SMSNumber: "+254700000001",
		},
		Stores: []StoreSpec{
			{
				ID:            "22222222-2222-2222-2222-222222222221",
				Name:          "Fresh Harvest Grocers",
				Location:      "Ground Floor",
				Participating: true,
				Products: []ProductSpec{
					{
						ID:          "33333333-3333-3333-3333-333333333311",
						Name:        "Bananas",
						Description: "Sweet ripe bananas sold per bunch.",
						SKU:         "DEMO-BANANAS-01",
						Price:       6,
					},
					{
						ID:          "33333333-3333-3333-3333-333333333312",
						Name:        "Spinach",
						Description: "Fresh spinach bundle for home cooking.",
						SKU:         "DEMO-SPINACH-01",
						Price:       4,
					},
				},
			},
			{
				ID:            "22222222-2222-2222-2222-222222222222",
				Name:          "Pantry Corner",
				Location:      "First Floor",
				Participating: true,
				Products: []ProductSpec{
					{
						ID:          "33333333-3333-3333-3333-333333333321",
						Name:        "Rice 2kg",
						Description: "Long-grain rice for family meals.",
						SKU:         "DEMO-RICE-2KG",
						Price:       18,
					},
					{
						ID:          "33333333-3333-3333-3333-333333333322",
						Name:        "Cooking Oil 1L",
						Description: "Sunflower cooking oil, 1 litre.",
						SKU:         "DEMO-OIL-1L",
						Price:       14,
					},
				},
			},
		},
	}
}
