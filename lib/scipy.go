package ris

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"gonum.org/v1/gonum/stat"
)

// FitRISParamsScipy calls the Python/Scipy implementation via subprocess
func FitRISParamsScipy(data []DataPoint, normalizer float64) (FitResult, error) {
	if len(data) < 5 {
		return FitResult{}, ErrNotEnoughData
	}

	// 1. Prepare data for Python
	x := make([]float64, len(data))
	y := make([]float64, len(data))
	for i, dp := range data {
		x[i] = dp.BodyWeight
		y[i] = dp.Total
	}

	inputJSON, _ := json.Marshal(map[string][]float64{"x": x, "y": y})

	_, filename, _, _ := runtime.Caller(0) // gets the current file path
	libDir := filepath.Dir(filename)
	scriptPath := filepath.Join(libDir, "fit_bridge.py")

	if _, err := os.Stat(scriptPath); err != nil {
		return FitResult{}, fmt.Errorf("script not found: %v", err)
	}

	// #nosec G204: scriptPath is safely constructed
	cmd := exec.Command("python3", scriptPath)
	cmd.Stdin = bytes.NewReader(inputJSON)

	var out bytes.Buffer
	var stderr bytes.Buffer // Add this buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr // Capture errors here

	if err := cmd.Run(); err != nil {
		// Now you will see the ACTUAL Python error (e.g., "ModuleNotFoundError")
		return FitResult{}, fmt.Errorf("python error: %v, stderr: %s", err, stderr.String())
	}

	// 3. Parse Python results
	var pyRes struct {
		A, K, B, V, Q float64
		Error         string `json:"error"`
	}
	if err := json.Unmarshal(out.Bytes(), &pyRes); err != nil {
		return FitResult{}, err
	}
	if pyRes.Error != "" {
		return FitResult{}, fmt.Errorf("scipy error: %s", pyRes.Error)
	}

	params := Params{A: pyRes.A, K: pyRes.K, B: pyRes.B, V: pyRes.V, Q: pyRes.Q}

	var sse float64
	for i := range x {
		diff := GeneralizedLogistic(x[i], params) - y[i]
		sse += diff * diff
	}
	rmse := math.Sqrt(sse / float64(len(x)))

	scores := make([]float64, len(data))
	for i := range data {
		// Using your existing Inverse method
		idx := params.Inverse(x[i], normalizer)
		scores[i] = idx * y[i]
	}

	slope, intercept := stat.LinearRegression(x, scores, nil, false)

	return FitResult{
		Params:        params,
		LineSlope:     slope,
		LineIntercept: intercept,
		RMSE:          rmse,
	}, nil
}
