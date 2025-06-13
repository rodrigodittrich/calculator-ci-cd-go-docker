package domain

type Operation interface {
	Calculate(a, b float64) (float64, error)
}
