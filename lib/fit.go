package rislib

import (
	"gonum.org/v1/gonum/optimize"
)

type DataPoint struct {
	Total      float64
	Bodyweight float64
}

// FitRISParams optimiert die Parameter A, K, Q, B, V auf Basis der Datenpunkte
func FitRISParams(data []DataPoint, initial RISParams) (RISParams, error) {
	problem := optimize.Problem{
		Func: func(x []float64) float64 {
			params := RISParams{x[0], x[1], x[2], x[3], x[4]}
			var sum float64
			for _, d := range data {
				predicted := RIS(d.Total, d.Bodyweight, params)
				relativeError := (predicted - d.Total) / d.Total // Relativer Fehler
				sum += relativeError * relativeError
			}
			return sum // Minimierung des quadratischen Fehlers
		},
	}

	settings := &optimize.Settings{
		FuncEvaluations: 1e5,
		MajorIterations: 1000,
	}

	method := &optimize.NelderMead{}

	result, err := optimize.Minimize(problem, []float64{
		initial.A, initial.K, initial.Q, initial.B, initial.V,
	}, settings, method)

	if err != nil {
		return RISParams{}, err
	}

	return RISParams{
		A: result.X[0],
		K: result.X[1],
		Q: result.X[2],
		B: result.X[3],
		V: result.X[4],
	}, nil
}
