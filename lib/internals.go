package ris

import (
	"errors"
	"math"
)

var (
	ErrNotEnoughData = errors.New("not enough data points to fit model")
)

func GeneralizedLogistic(x float64, p Params) float64 {
	return p.A + (p.K-p.A)/(1+p.Q*math.Exp(-p.B*(x-p.V)))
}

func (p Params) Inverse(x, normalizer float64) float64 {
	return normalizer / GeneralizedLogistic(x, p)
}
