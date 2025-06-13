package usecase

import "errors"

type Divide struct{}

func (d Divide) Calculate(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("divisão por zero")
	}
	return x / y, nil
}
