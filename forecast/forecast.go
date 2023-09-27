package arima

import (
	"github.com/georgethomas111/timeseries_forecasting/arima"
)

func Train(data []float64) (*Arima, error) {
	// Not sure what p is yet
	p := 1
	d := 0
	q := 0
	P := 0
	D := 1
	Q := 1
	m := 5

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

func (a *Arima) ForecastUpper(data []float64, futureVals int) ([]float64, error) {
	result := arima.ForeCastARIMA(data, futureVals, a.config)
	return result.GetForecastUpperConf(), nil
}

func (a *Arima) ForecastLower(data []float64, futureVals int) ([]float64, error) {
	result := arima.ForeCastARIMA(data, futureVals, a.config)
	return result.GetForecastLowerConf(), nil
}
