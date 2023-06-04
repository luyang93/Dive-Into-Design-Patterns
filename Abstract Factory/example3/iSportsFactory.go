package main

import "fmt"

const (
	AdidasBrand = iota
	NikeBrand
)

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(BrandType int) (ISportFactory, error) {
	switch BrandType {
	case AdidasBrand:
		return &Adidas{}, nil
	case NikeBrand:
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("Wrong brand type passed")
	}
}
