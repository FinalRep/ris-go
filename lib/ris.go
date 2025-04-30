package rislib

import "math"

type RISParams struct {
	A float64
	K float64
	Q float64
	B float64
	V float64
}

// RIS computes the Relative Index for Streetlifting
func RIS(total, bodyweight float64, params RISParams) float64 {
	denominator := params.A + (params.K-params.A)/(1+params.Q*math.Exp(-params.B*(bodyweight-params.V)))
	return (total * 100) / denominator
}
