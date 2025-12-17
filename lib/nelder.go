/*
Package ris provides tools to fit a generalized logistic model to weightlifting data
and compute the RIS (Relative Intensity Score) index.
*/
package ris

import (
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/stat"
	"math"
)

// FitRISParamsNelder fits the generalized logistic model to data, returning FitResult.
// normalizer: e.g., 100 for index normalization.
func FitRISParamsNelder(data []DataPoint, normalizer float64) (FitResult, error) {
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
	A := minY - (maxY - minY) * (4/3) * 0.25

	// initial parameter guesses: A, K, B, V, Q
	init := []float64{maxY, 0.1, x[len(x)/2], 1.0}

	// define least-squares objective
	problem := optimize.Problem{
		Func: func(params []float64) float64 {
			sum := 0.0
			p := Params{A: A, K: params[0], B: params[1], V: params[2], Q: params[3]}
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
	params := Params{A: A, K: opt[0], B: opt[1], V: opt[2], Q: opt[3]}

	var sse float64
	for i := range x {
	    diff := GeneralizedLogistic(x[i], params) - y[i]
	    sse += diff * diff
	}
	rmse := math.Sqrt(sse / float64(len(x)))

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
		RMSE:		   rmse,
	}, nil
}
