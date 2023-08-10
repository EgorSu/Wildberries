package pattern

import "errors"

type Product interface {
}

type Product1 struct {
	params ProductParams
}

func (p Product1) getName() string {
	return p.params.name
}

type Product2 struct {
	params ProductParams
}

func (p Product2) getName() string {
	return p.params.name
}

type ProductParams struct {
	name string
}

func getProduct(number int) (Product, error) {
	switch number {
	case 1:
		return &Product1{ProductParams{"first product"}}, nil
	case 2:
		return &Product2{ProductParams{"second product"}}, nil
	default:
		return nil, errors.New("product not found")
	}
}
