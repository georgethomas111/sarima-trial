package arima

import "testing"

func TestArima(t *testing.T) {
	increasingData := []float64{2, 4, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 32, 30, 28, 26, 24}
	a, err := Train(increasingData)
	if err != nil {
		t.Error("Error while training ", err.Error())
	}

	var expected []float64
	expected, err = a.Forecast(increasingData, 10)
	if err != nil {
		t.Error("Error while forecasting ", err.Error())
	}

	t.Log("Expected =", expected)
}
