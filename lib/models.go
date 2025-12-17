package ris

type DataPoint struct {
	BodyWeight float64
	Total      float64
}

type Params struct {
	A, K, B, V, Q float64
}

type FitResult struct {
	Params        Params
	LineSlope     float64
	LineIntercept float64
	RMSE          float64
}
