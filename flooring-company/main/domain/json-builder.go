package domain

import "time"


type TaxContainerBuilder struct {
	taxDataList []*TaxData
	doCreate    bool
}

func NewTaxContainerBuilder() *TaxContainerBuilder {
	return new(TaxContainerBuilder)
}

func (taxFileBuilder *TaxContainerBuilder) AddTaxData(taxDataArgs ...*TaxData) *TaxContainerBuilder {
	taxFileBuilder.taxDataList = taxDataArgs
	return taxFileBuilder
}

func (taxFileBuilder *TaxContainerBuilder) CreateOnBuild() *TaxContainerBuilder {
	taxFileBuilder.doCreate = true
	return taxFileBuilder
}

func (taxFileBuilder *TaxContainerBuilder) Build() (*TaxContainer, error) {
	taxFile := new(TaxContainer)
	taxFile.TaxDataList = taxFileBuilder.taxDataList
	if taxFileBuilder.doCreate == true {
		f,err :=  taxFile.createFile()
		if err != nil {
			return nil, err
		}
		f.Close()
	}

	return taxFile,nil
}

type OrderContainerBuilder struct {
	date time.Time
	orders []*Order
	doCreate bool
}

func NewOrderContainerBuilder() *OrderContainerBuilder {
	return new(OrderContainerBuilder)
}

func (o *OrderContainerBuilder) AddOrderData(orderArgs ...*Order) *OrderContainerBuilder {
	o.orders = orderArgs
	return o
}

func (o *OrderContainerBuilder) Date(t time.Time) *OrderContainerBuilder {
	o.date = t
	return o
}

func(o *OrderContainerBuilder) CreateOnBuild() *OrderContainerBuilder {
	o.doCreate = true
	return o
}

func(o *OrderContainerBuilder) Build() (*OrdersContainer, error) {
	orderFile := new(OrdersContainer)
	orderFile.Orders = o.orders
	orderFile.Date = o.date
	if o.doCreate {
		f, err := orderFile.createFile()
		if err != nil {
			return nil, err
		}
		f.Close()
	}
	return orderFile, nil
}
