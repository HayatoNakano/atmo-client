package co2client

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReading(t *testing.T) {
	type testcase struct {
		expectedErr string
		input       string
		values      values
	}
	re := regexp.MustCompile(`CO2=(?P<co2>\d+),HUM=(?P<hum>\d+\.\d+),TMP=(?P<tmp>-?\d+\.\d+)`)

	c := Client{re: re}

	for _, e := range []testcase{
		{"", "CO2=497,HUM=42.0,TMP=29.3", values{co2: 497, hum: 42.0, tmp: 29.3}},
		{"", "CO2=731,HUM=44.4,TMP=29.7", values{co2: 731, hum: 44.4, tmp: 29.7}},
		{"invalid format: ''", "", values{}},
		{"invalid format: 'STP'", "STP", values{}},
	} {
		a, err := c.read(e.input)
		if e.expectedErr != "" {
			assert.EqualError(t, err, e.expectedErr)
		} else {
			assert.Nil(t, err)
			assert.EqualValues(t, e.values, *a)
		}
	}
}

func TestCorrecting(t *testing.T) {
	type testcase struct {
		raw      values
		expected values
	}
	for _, e := range []testcase{
		{values{co2: 497, hum: 42.0, tmp: 29.3}, values{co2: 497, hum: 48.783404520484964, tmp: 24.8}},
		{values{co2: 731, hum: 44.4, tmp: 29.7}, values{co2: 731, hum: 51.44662921348314, tmp: 25.2}},
	} {
		a := correct(e.raw)
		assert.EqualValues(t, e.expected, *a)
	}
}
