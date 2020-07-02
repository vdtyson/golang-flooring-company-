package domain

type(
	ProductData struct {
		ProductType string `json:"productType"`
		CostPerSquareFoot float64 `json:"costPerSquareFoot"`
		LaborCostPerSquareFoot float64 `json:"laborCostPerSquareFoot"`
	}
	ProductsContainer struct {
		products []*ProductData
	}
)
