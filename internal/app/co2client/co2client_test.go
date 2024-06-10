package co2client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReading(t *testing.T) {
	type values struct {
		co2 float64
		hum float64
		tmp float64
	}

	type testcase struct {
		input  string
		values values
	}

	for _, e := range []testcase{
		{"CO2=497,HUM=42.0,TMP=29.3", values{co2: 497, hum: 42.0, tmp: 29.3}},
		{"CO2=731,HUM=44.4,TMP=29.7", values{co2: 731, hum: 44.4, tmp: 29.7}},
	} {
		co2, hum, tmp, err := read(e.input)
		a := values{co2, hum, tmp}
		assert.Nil(t, err)
		assert.EqualValues(t, e.values, a)
	}
}
