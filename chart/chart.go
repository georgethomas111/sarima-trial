package chart

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"

	arima "github.com/georgethomas111/sarima-trial/forecast"
)

type Chart struct {
	Title    string
	XLabel   string
	YLabel   string
	X        []string
	Y        []float64
	YHalf    []float64
	YProject []float64
}

func New(title string, d []byte) (*Chart, error) {
	c := &Chart{
		Title:    title,
		XLabel:   "",
		YLabel:   "",
		X:        []string{},
		Y:        []float64{},
		YHalf:    []float64{},
		YProject: []float64{},
	}

	r := csv.NewReader(strings.NewReader(string(d)))
	labels := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if labels {
			c.XLabel = record[0]
			c.YLabel = record[1]
			labels = false
			continue
		}
		if err != nil {
			return nil, err
		}
		c.X = append(c.X, record[0])

		yFloat, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}
		c.Y = append(c.Y, yFloat)
	}

	yLen := len(c.Y)
	c.YHalf = make([]float64, yLen/2)
	copy(c.YHalf, c.Y)

	a, err := arima.Train(c.YHalf)
	if err != nil {
		return nil, err
	}

	var forecasts []float64
	forecasts, err = a.Forecast(c.YHalf, yLen/2)
	if err != nil {
		return nil, err
	}

	c.YProject = append(c.YHalf, forecasts...)
	return c, nil
}

// Render renders the chart based on the input parameters
func (c *Chart) Render() {

	t, err := template.New("webpage").Parse(ChartTmp)
	err = t.Execute(os.Stdout, c)
	if err != nil {
		fmt.Println(" Error =", err.Error())
	}
}
