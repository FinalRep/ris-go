/*
Package ris provides tools to fit a generalized logistic model to weightlifting data
and compute the RIS (Relative Intensity Score) index.
*/
package ris

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/stat"
)

var (
	ErrNotEnoughData = errors.New("not enough data points to fit model")
)

// DataPoint holds one observation: bodyweight and total lift
type DataPoint struct {
	BodyWeight float64 // BW
	Total      float64 // Total (kg)
}

// Params are the generalized logistic parameters
// A: lower asymptote
// K: upper asymptote
// B: growth rate
// V: midpoint location
// Q: related to value at x=0
type Params struct {
	A, K, B, V, Q float64
}

// FitResult holds the fitted Params and resulting linear approximation
type FitResult struct {
	Params        Params  // fitted model parameters
	LineSlope     float64 // slope for RIS*Total vs BW
	LineIntercept float64 // intercept for RIS*Total vs BW
}

// GeneralizedLogistic computes the value of the model at x
func GeneralizedLogistic(x float64, p Params) float64 {
	return p.A + (p.K-p.A)/(1+p.Q*math.Exp(-p.B*(x-p.V)))
}

// Inverse computes normalizer / logistic(x)
func (p Params) Inverse(x, normalizer float64) float64 {
	return normalizer / GeneralizedLogistic(x, p)
}

// FitRISParams fits the generalized logistic model to data, returning FitResult.
// normalizer: e.g., 100 for index normalization.
func FitRISParams(data []DataPoint, normalizer float64) (FitResult, error) {
	if len(data) < 5 {
		return FitResult{}, ErrNotEnoughData
	}

	// prepare x,y slices
	x := make([]float64, len(data))
	y := make([]float64, len(data))
	for i, dp := range data {
		x[i] = dp.BodyWeight
		y[i] = dp.Total
	}

	// estimate A: min(y) - (max(y)-min(y))/3
	minY, maxY := floats.Min(y), floats.Max(y)
	Ainit := minY - (maxY-minY)/3.0

	// initial parameter guesses: A, K, B, V, Q
	init := []float64{Ainit, maxY, 0.01, x[len(x)/2], 1.0}

	// define least-squares objective
	problem := optimize.Problem{
		Func: func(params []float64) float64 {
			sum := 0.0
			p := Params{A: params[0], K: params[1], B: params[2], V: params[3], Q: params[4]}
			for i := range x {
				d := GeneralizedLogistic(x[i], p) - y[i]
				sum += d * d
			}
			return sum
		},
	}

	// perform optimization using Nelder-Mead
	settings := optimize.Settings{GradientThreshold: 1e-8}
	method := &optimize.NelderMead{}
	result, err := optimize.Minimize(problem, init, &settings, method)
	if err != nil {
		return FitResult{}, err
	}

	opt := result.X
	params := Params{A: opt[0], K: opt[1], B: opt[2], V: opt[3], Q: opt[4]}

	// compute RIS*Total for each data point and fit linear model
	scores := make([]float64, len(data))
	for i := range data {
		idx := params.Inverse(x[i], normalizer)
		scores[i] = idx * y[i]
	}
	// linear regression: scores ~ x
	slope, intercept := stat.LinearRegression(x, scores, nil, false)

	return FitResult{
		Params:        params,
		LineSlope:     slope,
		LineIntercept: intercept,
	}, nil
}
