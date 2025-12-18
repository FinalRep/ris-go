package ris

// DataPoint contains the dataset information to calculate our fitting
type DataPoint struct {
	BodyWeight float64
	Total      float64
}

// Params contains our logistic function parameters
type Params struct {
	A, K, B, V, Q float64
}

// FitResult contains our optimization result and additional meta information
type FitResult struct {
	Params        Params
	LineSlope     float64
	LineIntercept float64
	RMSE          float64
}
