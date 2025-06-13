package usecase

type Multiply struct{}

func (m Multiply) Calculate(x, y float64) (float64, error) {
	return x * y, nil
}
