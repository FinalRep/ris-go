package ris

func fittingIsInRange(fit Params, male bool) bool {
	if male {
		if fit.A < 300 && fit.A > 350 {
			return false
		}
		if fit.B < 0 && fit.B > 1 {
			return false
		}
		if fit.K < 500 && fit.K > 550 {
			return false
		}
		if fit.Q < 0 && fit.Q > 20 {
			return false
		}
		if fit.V < 200 && fit.V > 250 {
			return false
		}
		return true
	}

	// female
	if fit.A < 150 && fit.A > 250 {
		return false
	}
	if fit.B < 0 && fit.B > 1 {
		return false
	}
	if fit.K < 200 && fit.K > 300 {
		return false
	}
	if fit.Q < 0 && fit.Q > 20 {
		return false
	}
	if fit.V < 50 && fit.V > 100 {
		return false
	}
	return true
}
