package arima

import (
	"github.com/georgethomas111/timeseries_forecasting/arima"
)

func Train(data []float64) (*Arima, error) {
	// Not sure what p is yet
	p := 4
	d := 1
	q := 2
	P := 1
	D := 1
	Q := 0
	m := 0

	c := arima.NewConfig(p, d, q, P, D, Q, m)

	return &Arima{
		config: c,
	}, nil
}

type Arima struct {
	config arima.Config
}

func (a *Arima) Forecast(data []float64, futureVals int) ([]float64, error) {
	result := arima.ForeCastARIMA(data, futureVals, a.config)
	return result.GetForecast(), nil
}
