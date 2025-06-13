package usecase

type Subtract struct{}

func (s Subtract) Calculate(x, y float64) (float64, error) {
	return x - y, nil
}
