package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Fachrurazzi",
		MiddleName: "Razzy",
		LastName:   "Tirta",
	}

	encoder.Encode(customer)
}
