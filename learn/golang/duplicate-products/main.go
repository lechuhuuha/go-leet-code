package main

import (
	"encoding/json"
	"fmt"
)

/*
QUESTION: Duplicate Products

You are given a list of products. Each product contains the following attributes:
  - Code (int)
  - Organisation (string)
  - Name (string)

Write a function to list all duplicated products. In the current business context,
a product is identified as a duplicate if the combination of Code and Organisation
is the same with another product that appeared earlier in the list.
Input and output will be a list of product structs (pointers).

Sample Input:
  products := []*Product{
      {Code: 1, Organisation: "Tech Corp", Name: "SuperWidget"},
      {Code: 2, Organisation: "Gadget Inc", Name: "MegaGadget"},
      {Code: 3, Organisation: "Device LLC", Name: "UltraDevice"},
      {Code: 1, Organisation: "Tech Corp", Name: "HyperWidget"},
      {Code: 2, Organisation: "Gadget Inc", Name: "GigaGadget"},
      {Code: 1, Organisation: "Tech Corp", Name: "HyperWidget"},
  }

Sample Output:
  [
      {
          "Code": 1,
          "Organisation": "Tech Corp",
          "Name": "HyperWidget"
      },
      {
          "Code": 2,
          "Organisation": "Gadget Inc",
          "Name": "GigaGadget"
      },
      {
          "Code": 1,
          "Organisation": "Tech Corp",
          "Name": "HyperWidget"
      }
  ]

Follow-up:
  - What is the performance of your solution in Big-O complexity (Time and Space)?
  - Time is O(n) since i need to traverse all the slice
  - Space is O(1) since i only access key
Commands:
  Run program:   go run .
  Run tests:     go test -v .
*/

// Product represents a product with Code, Organisation, and Name.
type Product struct {
	Code         int    `json:"Code"`
	Organisation string `json:"Organisation"`
	Name         string `json:"Name"`
}

// FindDuplicateProducts returns all products identified as duplicates.
// A product is identified as a duplicate if its combination of Code and Organisation
// has already appeared earlier in the list.
func FindDuplicateProducts(products []*Product) []*Product {
	// TODO: Implement your solution here
	dupProducts := make(map[string]struct{})
	listProductsDup := make([]*Product, 0)
	for _, p := range products {
		key := fmt.Sprintf("%d|%s", p.Code, p.Organisation)
		if _, ok := dupProducts[key]; ok {
			listProductsDup = append(listProductsDup, p)
		} else {
			dupProducts[key] = struct{}{}
		}
	}

	return listProductsDup
}

func main() {
	products := []*Product{
		{Code: 1, Organisation: "Tech Corp", Name: "SuperWidget"},
		{Code: 2, Organisation: "Gadget Inc", Name: "MegaGadget"},
		{Code: 3, Organisation: "Device LLC", Name: "UltraDevice"},
		{Code: 1, Organisation: "Tech Corp", Name: "HyperWidget"},
		{Code: 2, Organisation: "Gadget Inc", Name: "GigaGadget"},
		{Code: 1, Organisation: "Tech Corp", Name: "HyperWidget"},
	}

	duplicates := FindDuplicateProducts(products)

	output, err := json.MarshalIndent(duplicates, "", "    ")
	if err != nil {
		fmt.Printf("Error marshalling output: %v\n", err)
		return
	}

	fmt.Println("Duplicated Products:")
	fmt.Println(string(output))
}
