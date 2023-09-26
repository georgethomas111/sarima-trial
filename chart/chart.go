package chart

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"
)

type Chart struct {
	Title  string
	XLabel string
	YLabel string
	X      []string
	Y      []float64
}

func New(title string, d []byte) (*Chart, error) {
	c := &Chart{
		Title:  title,
		XLabel: "",
		YLabel: "",
		X:      []string{},
		Y:      []float64{},
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
