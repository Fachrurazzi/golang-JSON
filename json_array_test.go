package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Fachrurazzi",
		MiddleName: "Razzy",
		LastName:   "Tirta",
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fachrurazzi","MiddleName":"Razzy","LastName":"Tirta","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Razzi",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "99999",
			},
			{
				Street:     "Jalan Belum Jadi",
				Country:    "Indonesia",
				PostalCode: "88888",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Razzi","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"99999"},{"Street":"Jalan Belum Jadi","Country":"Indonesia","PostalCode":"88888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonArray := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"99999"},{"Street":"Jalan Belum Jadi","Country":"Indonesia","PostalCode":"88888"}]`
	jsonBytes := []byte(jsonArray)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "99999",
		},
		{
			Street:     "Jalan Belum Jadi",
			Country:    "Indonesia",
			PostalCode: "88888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
