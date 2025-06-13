package usecase

type Add struct{}

func (a Add) Calculate(x, y float64) (float64, error) {
	return x + y, nil
}
