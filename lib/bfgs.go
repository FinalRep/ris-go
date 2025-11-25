package ris

import (
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/stat"
)

func FitRISParamsBFGS(data []DataPoint, normalizer float64) (FitResult, error) {
	if len(data) < 5 {
		return FitResult{}, ErrNotEnoughData
	}

	n := len(data)
	x := make([]float64, n)
	y := make([]float64, n)
	for i, dp := range data {
		x[i] = dp.BodyWeight
		y[i] = dp.Total
	}

	// A fixed based on your heuristic
	minY, maxY := floats.Min(y), floats.Max(y)
	A := minY - (maxY-minY)/3.0

	// Initial guess: K, B, V, Q
	init := []float64{maxY, 0.01, x[n/2], 1.0}

	// Define optimization problem: minimize sum of squared residuals
	problem := optimize.Problem{

		Func: func(pv []float64) float64 {
			p := Params{A: A, K: pv[0], B: pv[1], V: pv[2], Q: pv[3]}
			sum := 0.0
			for i := range x {
				diff := GeneralizedLogistic(x[i], p) - y[i]
				sum += diff * diff
			}
			return sum
		},

		Grad: func(pv, grad []float64) {
			p := Params{A: A, K: pv[0], B: pv[1], V: pv[2], Q: pv[3]}

			// Zero gradient
			for i := range grad {
				grad[i] = 0
			}

			for i := range x {
				xi := x[i]
				yi := y[i]

				// intermediate terms
				z := math.Exp(-p.B * (xi - p.V)) // exp(-B*(x-V))
				d := 1 + p.Q*z                   // denominator
				model := p.A + (p.K-p.A)/d
				resid := model - yi

				// partial derivs of model
				// d(model)/dK
				dK := 1.0 / d

				// d(model)/dB
				dB := (p.K - p.A) * p.Q * z * (xi - p.V) / (d * d)

				// d(model)/dV
				dV := (p.K - p.A) * p.Q * z * p.B / (d * d)

				// d(model)/dQ
				dQ := -(p.K - p.A) * z / (d * d)

				grad[0] += 2 * resid * dK
				grad[1] += 2 * resid * dB
				grad[2] += 2 * resid * dV
				grad[3] += 2 * resid * dQ
			}
		},
	}

	settings := optimize.Settings{
		GradientThreshold: 1e-8,
		MajorIterations:   1000,
	}

	method := &optimize.BFGS{}

	// Run optimizer
	result, err := optimize.Minimize(problem, init, &settings, method)
	if err != nil {
		return FitResult{}, err
	}

	// Extract fitted parameters
	opt := result.X
	params := Params{
		A: A,
		K: opt[0],
		B: opt[1],
		V: opt[2],
		Q: opt[3],
	}

	// Compute RIS*Total for each data point → then fit linear model
	scores := make([]float64, n)
	for i := range data {
		idx := params.Inverse(x[i], normalizer)
		scores[i] = idx * y[i]
	}

	slope, intercept := stat.LinearRegression(x, scores, nil, false)

	return FitResult{
		Params:        params,
		LineSlope:     slope,
		LineIntercept: intercept,
	}, nil
}
