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

type PrincipalSpec struct {
	ID         string
	Name       string
	Email      string
	Password   string
	Avatar     string
	Status     string
	Roles      []string
	CustomerID string
	Kind       string
}

type BootstrapSpec struct {
	Customer CustomerSpec
	Auth     PrincipalSpec
	BotAuth  PrincipalSpec
	Stores   []StoreSpec
}

func Spec() BootstrapSpec {
	return BootstrapSpec{
		Customer: CustomerSpec{
			ID:        "11111111-1111-1111-1111-111111111111",
			Name:      "Demo Shopper",
			SMSNumber: "+254700000001",
		},
		Auth: PrincipalSpec{
			ID:         "cfaad35d-07a3-4447-a6c3-d8c3d54fd5df",
			Name:       "Brian Hughes",
			Email:      "hughes.brian@company.com",
			Password:   "admin",
			Avatar:     "images/avatars/brian-hughes.jpg",
			Status:     "online",
			Roles:      []string{"admin", "customer"},
			CustomerID: "11111111-1111-1111-1111-111111111111",
			Kind:       "user",
		},
		BotAuth: PrincipalSpec{
			ID:       "ae1f81af-e4a0-4fa7-90d4-b0f1dad3c001",
			Name:     "MallBots Service Bot",
			Email:    "bot@mallbots.internal",
			Password: "",
			Avatar:   "",
			Status:   "online",
			Roles:    []string{"bot"},
			Kind:     "bot",
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
