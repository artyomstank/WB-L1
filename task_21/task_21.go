package main

import "fmt"

type TempTool interface {
	getCurrentTemp() float64
}

type OldTempTool struct {
	tempInFahrenheit float64
}

func (cts *OldTempTool) getTemp() float64 {
	return cts.tempInFahrenheit
}

type AdapterTempTool struct {
	cts *OldTempTool
}

func (att AdapterTempTool) getCurrentTemp() float64 {
	return (att.cts.tempInFahrenheit - 32.0) * 5.0 / 9.0
}

var cts = OldTempTool{tempInFahrenheit: 150}

func main() {
	var att TempTool = AdapterTempTool{cts: &cts}

	fmt.Printf("Температура в С: %f", att.getCurrentTemp())
}
