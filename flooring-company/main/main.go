package main

import (
	"mthree/flooring-company/main/domain"
)


var orderFiles = make([]*domain.OrdersContainer, 0)

func appendOrderFile(orderFile *domain.OrdersContainer) {
	orderFiles = append(orderFiles, orderFile)
}
func main() {
}


/*func CreateProductFile() *os.File {
	orderFile, err := os.create(MainDirectory + ProductFilePath)
	if err != nil {
		panic(err)
	}
}*/

// bw.WriteString(fmt.Sprintf("%s,%s,%s", State, usc[State], TaxRate))
