package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "amits",
		Price: 1.00,
		SKU:   "abs-abd-efg-ers",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
