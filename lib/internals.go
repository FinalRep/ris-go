package ris

import (
	"errors"
	"math"
)

var (
	// ErrNotEnoughData if we lack on the amount of data to make a proper fitting
	ErrNotEnoughData = errors.New("not enough data points to fit model")
)

// GeneralizedLogistic function calculates the logistic function using given parameters p
func GeneralizedLogistic(x float64, p Params) float64 {
	return p.A + (p.K-p.A)/(1+p.Q*math.Exp(-p.B*(x-p.V)))
}

// Inverse inverts our logistic function
func (p Params) Inverse(x, normalizer float64) float64 {
	return normalizer / GeneralizedLogistic(x, p)
}
