package ris

import "math"

// RIS computes the Relative Index for Streetlifting
func RIS(total, bodyweight float64, params Params) float64 {
	denominator := params.A + (params.K-params.A)/(1+params.Q*math.Exp(-params.B*(bodyweight-params.V)))
	return (total * 100) / denominator
}
