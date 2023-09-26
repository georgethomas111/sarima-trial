package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/georgethomas111/sarima-trial/chart"
)

func main() {
	csvInput := os.Args[1]

	d, err := ioutil.ReadFile(csvInput)
	if err != nil {
		fmt.Println("Could not read csvFile passed ", csvInput, " error = ", err.Error())
		return
	}

	var c *chart.Chart

	title := strings.Split(csvInput, ".")[0]
	c, err = chart.New(title, d)
	if err != nil {
		fmt.Println("Could not initialize chart ", err.Error())
		return
	}
	c.Render()
}
