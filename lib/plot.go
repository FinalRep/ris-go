package ris

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// PlotFitGraph plots the datapoints and the fitted generalized logistic curve.
//nolint:funlen // TBD
func PlotFitGraph(data []DataPoint, result FitResult, title, filename string) error {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "Bodyweight (kg)"
	p.Y.Label.Text = "Total (kg)"
	p.Add(plotter.NewGrid())

	xys := make(plotter.XYs, len(data))
	sizes := make(plotter.XYs, len(data))
	colors := make([]color.RGBA, len(data))

	scatterData := make(plotter.XYs, len(data))
	for i, dp := range data {
		scatterData[i].X = dp.BodyWeight
		scatterData[i].Y = dp.Total
	}

	maxResidual := 0.0
	for _, d := range data {
		fitY := GeneralizedLogistic(d.BodyWeight, result.Params)
		residual := math.Abs(fitY - d.Total)
		if residual > maxResidual {
			maxResidual = residual
		}
	}

	for i, d := range data {
		xys[i].X = d.BodyWeight
		xys[i].Y = d.Total

		fitY := GeneralizedLogistic(d.BodyWeight, result.Params)
		residual := math.Abs(fitY - d.Total)
		normTotal := (d.Total - 350) / (650 - 350) // Normalize to 0–1 range
		size := 3.0 + 8.0*math.Pow(normTotal, 2)   // Quadratic scaling for better spread
		sizes[i].X = size
		sizes[i].Y = size

		ratio := residual / maxResidual
		red := uint8(255)
		green := uint8(ratio * 255)
		colors[i] = color.RGBA{R: red, G: green, B: 0, A: 255}
	}

	// Border layer
	borderScatter, err := plotter.NewScatter(xys)
	if err != nil {
		return err
	}
	borderScatter.GlyphStyleFunc = func(i int) draw.GlyphStyle {
		return draw.GlyphStyle{
			Color:  color.RGBA{R: 0, G: 0, B: 0, A: 100},
			Radius: vg.Length(sizes[i].X + 1.0), // Slightly larger
			Shape:  draw.CircleGlyph{},
		}
	}
	p.Add(borderScatter)

	// Colored data layer
	scatter, err := plotter.NewScatter(xys)
	if err != nil {
		return err
	}
	scatter.GlyphStyleFunc = func(i int) draw.GlyphStyle {
		return draw.GlyphStyle{
			Color:  colors[i],
			Radius: vg.Length(sizes[i].X),
			Shape:  draw.CircleGlyph{},
		}
	}
	p.Add(scatter)

	fitLine := make(plotter.XYs, 200)
	minX, maxX := scatterData[0].X, scatterData[0].X
	minY, maxY := scatterData[0].Y, scatterData[0].Y
	for _, pt := range scatterData {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}
	minX = minX - 10
	maxX = maxX + 10
	p.X.Tick.Marker = plot.ConstantTicks(generateTicks(minX, maxX, 5))
	p.Y.Tick.Marker = plot.ConstantTicks(generateTicks(minY, maxY, 50))

	step := (maxX - minX) / float64(len(fitLine)-1)
	for i := range fitLine {
		x := minX + float64(i)*step
		fitLine[i].X = x
		fitLine[i].Y = GeneralizedLogistic(x, result.Params)
	}

	line, err := plotter.NewLine(fitLine)
	if err != nil {
		return err
	}
	line.Color = color.RGBA{R: 255, A: 255}
	line.Width = vg.Points(1.5)

	p.Add(scatter, line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		log.Printf("Failed to save plot: %v", err)
		return err
	}
	return nil
}

func generateTicks(minimum, maximum, step float64) []plot.Tick {
	var ticks []plot.Tick
	for val := math.Ceil(minimum/step) * step; val <= maximum; val += step {
		ticks = append(ticks, plot.Tick{Value: val, Label: fmt.Sprintf("%.0f", val)})
	}
	return ticks
}
