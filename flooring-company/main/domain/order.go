package domain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	// MMDDYYYY
	DateOutput             = "01022006"
	JSONPrefix             = ""
	JSONIndent             = "    "
	OrderNumber            = "OrderNumber"
	CustomerName           = "CustomerName"
	State                  = "State"
	TaxRate                = "TaxRate"
	ProductType            = "ProductType"
	Area                   = "Area"
	CostPerSquareFoot      = "CostPerSquareFoot"
	LaborCostPerSquareFoot = "LaborCostPerSquareFoot"
	MaterialCost           = "MaterialCost"
	LaborCost              = "LaborCost"
	Tax                    = "Tax"
	Total                  = "Total"
)

type (
	Order struct {
		OrderNumber            int     `json:"orderNumber"`
		CustomerName           string  `json:"customerName"`
		State                  string  `json:"stateAbbreviation"`
		TaxRate                float64 `json:"taxRate"`
		ProductType            string  `json:"productType"`
		Area                   float64 `json:"area"`
		CostPerSquareFoot      float64 `json:"costPerSquareFoot"`
		LaborCostPerSquareFoot float64 `json:"laborCostPerSquareFoot"`
		MaterialCost           float64 `json:"materialCost"`
		LaborCost              float64 `json:"laborCost"`
		Tax                    float64 `json:"tax"`
		Total                  float64 `json:"total"`
	}

	OrdersContainer struct {
		Date      time.Time `json:"Date"`
		Orders    []*Order  `json:"Orders"`
		isCreated bool
	}
)

func (ordersContainer *OrdersContainer) WriteJson() error {
	file, err := ordersContainer.GetOrCreateFile()
	if err != nil {
		return err
	}
	bw := bufio.NewWriter(file)
	data, err := json.MarshalIndent(ordersContainer, JSONPrefix, JSONIndent)
	if err != nil {
		return err
	}
	_, err = bw.Write(data)
	if err != nil {
		return err
	}
	err = bw.Flush()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ordersContainer *OrdersContainer) WasFileCreated() bool {
	return ordersContainer.isCreated
}

func (ordersContainer *OrdersContainer) FilePath() string {
	dateOutput := ordersContainer.Date.Format(DateOutput)

	return fmt.Sprintf(`%sorders\Orders_%s.json`, MainDirectory, dateOutput)
}

func (ordersContainer *OrdersContainer) createFile() (*os.File, error) {
	if ordersContainer.isCreated {
		return nil, fmt.Errorf("file already created")
	}
	jsonFile, err := os.Create(ordersContainer.FilePath())
	if err != nil {
		return nil, err
	}
	ordersContainer.isCreated = true
	return jsonFile, nil
}

func (ordersContainer *OrdersContainer) RemoveFile() error {
	err := os.Remove(ordersContainer.FilePath())
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (ordersContainer *OrdersContainer) openAndGetFile() (*os.File, error) {
	return os.OpenFile(ordersContainer.FilePath(), os.O_RDWR, 0666)
}

func (ordersContainer *OrdersContainer) GetOrCreateFile() (*os.File, error) {
	if ordersContainer.isCreated {
		return ordersContainer.openAndGetFile()
	} else {
		return ordersContainer.createFile()
	}
}
