package main

//cd C:\Users\Kali69\go\src\github.com\Rob8150\weather>
//go tests
//go test -vet=off
//actual path
//C:\Users\Kali69\go\src\github\Rob8150\weather\weather.go

//cd C:\Users\Kali69\go
//go install github\Rob8150\weather
//weather

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestWeather(t *testing.T) {

	type testScenario = struct {
		T   float64 `json:"T"`
		RH  float64 `json:"Rh"`
		Mb  float64 `json:"Mb"`
		Alt float64 `json:"Alt"`
		E   float64 `json:"E"`
		Es  float64 `json:"Es"`
		Tdc float64 `json:"Tdc"`
		Ah  float64 `json:"Ah"`
		SH  float64 `json:"SH"`
		Cb  float64 `json:"Cb"`
	}

	var TestScenarios []testScenario
	//first way of creating tests
	TestScenarios = append(TestScenarios, testScenario{T: 27.8, RH: 40, Mb: 1016.8, Alt: 0, E: 14.889921514085886, Es: 37.269360833764715, Tdc: 12.927401408783576, Ah: 10.720423078136879, SH: 9.159208084179944, Cb: 1859.0748239020531})
	TestScenarios = append(TestScenarios, testScenario{T: 32, RH: 28, Mb: 1013.25, Alt: 230, E: 13.250897568572489, Es: 47.41533814132679, Tdc: 11.155498055449069, Ah: 9.409054524775867, SH: 8.174689358817522, Cb: 2835.5627430688664})
	TestScenarios = append(TestScenarios, testScenario{T: 40, RH: 55, Mb: 1013.25, Alt: 0, E: 40.985853037166194, Es: 73.50953828383807, Tdc: 29.442396392805712, Ah: 28.359331466575703, SH: 25.550501483818934, Cb: 1319.700450899286})
	//second way of creating test
	jso := `{"T":40,"Rh":55,"Mb":1013.25,"Alt":0,"E":40.985853037166194,"Es":73.50953828383807,"Tdc":29.442396392805712,"Ah":28.359331466575703,"SH":25.550501483818934,"Cb":1319.700450899286}`
	TestScenarios = append(TestScenarios, jsonToObj(jso))

	jso = `{"T":18,"Rh":75,"Mb":1013.25,"Alt":0,"E":15.523379594843481,"Es":20.60647308856927,"Tdc":13.566553520797422,"Ah":11.552683125489954,"SH":9.584785662577794,"Cb":554.1808099003223}`
	TestScenarios = append(TestScenarios, jsonToObj(jso))

	jso = `{"T":5,"Rh":95,"Mb":1016.25,"Alt":130,"E":8.270629434587521,"Es":8.720732259078778,"Tdc":4.241963689953126,"Ah":6.442763352321696,"SH":5.077693373261027,"Cb":224.75453875585924}`
	TestScenarios = append(TestScenarios, jsonToObj(jso))

	for i, scenario := range TestScenarios {
		fmt.Println("testing", i)
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			expected := scenario.Es
			actual := VapourSat(scenario.T)
			if actual != expected {
				t.Errorf("%g %g VapourSat() Test Failed", expected, actual)
			}

			expected2 := scenario.E
			actual2 := VapourPartial(scenario.T, scenario.RH)
			if actual2 != expected2 {
				t.Errorf("%g %g VapourPartial() Test Failed", expected2, actual2)
			}

			expected3 := scenario.Tdc
			actual3 := DewPoint(scenario.T, scenario.RH)
			if actual3 != expected3 {
				t.Errorf("%g %g DewPoint() Test Failed", expected3, actual3)
			}

			expected4 := scenario.Cb
			actual4 := CloudBase(scenario.T, actual3, scenario.Alt)
			if actual4 != expected4 {
				t.Errorf("%g %g CloudBase() Test Failed", expected4, actual4)
			}

			expected5 := scenario.Ah
			actual5 := AbsoluteHum(scenario.T, actual3)
			if actual5 != expected5 {
				t.Errorf("%g %g AbsoluteHum() Test Failed", expected5, actual5)
			}

			expected6 := scenario.SH
			actual6 := SpecificHum(scenario.Mb, actual2)
			if actual6 != expected6 {
				t.Errorf("%g %g SpecificHum() Test Failed", expected6, actual6)
			}

		})
	}
}

func jsonToObj(jso string) testScenario {
	b := []byte(jso)

	var m testScenario
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Err", err)
	}
	return m
}
