package get_cocoa

import (
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

type PlotStruct struct {
	X float64
	Y float64
}

func Plot(series plotter.XYs) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Time Series"
	p.X.Label.Text = "Day #"
	p.Y.Label.Text = "Price"

	err = plotutil.AddLinePoints(p,
		"Close", series,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Save(12*vg.Inch, 6*vg.Inch, "price.png"); err != nil {
		log.Fatal(err)
	}
}
