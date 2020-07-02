package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mthree/flooring-company/main/domain"
	"os"
	"testing"
	"time"
)

var testDate = time.Date(2020, 8, 14, 0, 0, 0, 0, time.UTC)
var testOrders = []*domain.Order {
	{322, "John", "IN", 5.33, "Sumn", 322.1, 322.2, 3222, 322.2, 322.2, 322.2, 2232.2},
	{382, "Jack", "AR", 5.33, "Sumn", 322.1, 322.2, 3222, 322.2, 322.2, 322.2, 2232.2},
	{392, "Randal", "CA", 5.33, "Sumn", 322.1, 322.2, 3222, 322.2, 322.2, 322.2, 2232.2},
}

// test
func TestMainDirectory(t *testing.T) {
	// GIVEN
	filePath := "test.txt"
	// WHEN
	f, err := os.Create(domain.MainDirectory + filePath)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = f.Close()
	if err != nil {
		t.Error(err.Error())
	}
	// THEN
	err = os.Remove(domain.MainDirectory + filePath)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestOrderContainerBuilder(t *testing.T) {
	// Given
	orders := []*domain.Order {
		{322, "John", "IN", 5.33, "Sumn", 322.1, 322.2, 3222, 322.2, 322.2, 322.2, 2232.2},
		{382, "Jack", "AR", 5.33, "Sumn", 322.1, 322.2, 3222, 322.2, 322.2, 322.2, 2232.2},
		{392, "Randal", "CA", 5.33, "Sumn", 322.1, 322.2, 3222, 322.2, 322.2, 322.2, 2232.2},
	}
	builder := domain.NewOrderContainerBuilder()
	var orderContainer *domain.OrdersContainer
	// When
	orderContainer, err := builder.Date(testDate).AddOrderData(orders...).Build()
	if err != nil {
		t.Errorf(err.Error())
	}
	// Then
	if fmt.Sprintf("%v", orders) != fmt.Sprintf("%v", orderContainer.Orders) {
		t.Errorf("Orders are not the same\n")
	}
	if testDate != orderContainer.Date {
		t.Errorf("Order container date mismatch. Expected: %s, Actual: %s\n", testDate, orderContainer.Date)
	}
}

func TestTaxContainerBuilder(t *testing.T) {
	// GIVEN
	taxDataList := []*domain.TaxData {
		{"IN", "Indiana", 3.45},
		{"AR", "Arizona", 8.99},
		{"MI", "Michigan", 4.55},
	}
	builder := domain.NewTaxContainerBuilder()
	var taxContainer *domain.TaxContainer
	var expected = fmt.Sprintf("%v", taxDataList)
	var actual string
	// WHEN
	taxContainer, err := builder.AddTaxData(taxDataList...).Build()
	if err != nil {
		t.Errorf(err.Error())
	}
	actual = fmt.Sprintf("%v", taxContainer.TaxDataList)
	// THEN
	if expected != actual {
		t.Errorf("Tax Data mismatch. Expected: %s, Actual: %s\n", expected, actual)
	}
}


func TestDateOutputFormat(t *testing.T) {
	// GIVEN
	date := time.Date(2020, 8, 14, 0, 0, 0, 0, time.UTC)
	expected := "08142020"
	var actual string
	// WHEN
	actual = date.Format(domain.DateOutput)
	// THEN
	if actual != expected {
		t.Errorf("Date output incorrect. actual: %s expected: %s\n", actual, expected)

	}
}

func TestOrderContainerMethodWriteJson(t *testing.T) {
	// GIVEN
	orderContainer, err := domain.NewOrderContainerBuilder().Date(testDate).AddOrderData(testOrders...).CreateOnBuild().Build()
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(orderContainer, domain.JSONPrefix, domain.JSONIndent)
	if err != nil {
		panic(err)
	}
	expected := string(data)
	var actual string
	// WHEN
	err = orderContainer.WriteJson()
	if err != nil {
		t.Error(err.Error())
	}

	bytes, err := ioutil.ReadFile(orderContainer.FilePath())
	if err != nil {
		t.Error(err.Error())
	}
	actual = string(bytes)
	// THEN
	if expected != actual {
		t.Errorf("Json mismatch: Actual: %s, Expected: %s", actual, expected)
	}
	err = os.Remove(orderContainer.FilePath())
	if err != nil {
		t.Error(err.Error())
	}
}
