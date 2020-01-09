package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_GetTLVValue(t *testing.T) {
	var tests = []struct {
		caseName string
		str      string
		size     int
		want     string
		err      error
	}{
		{
			"Success",
			"A0511",
			3,
			"A05",
			nil,
		},
		{
			"Error",
			"A0511",
			8,
			"",
			errors.New("string require min 8 chars"),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.caseName)
		t.Run(testname, func(t *testing.T) {
			subStr, err := getTLVValue(tt.str, tt.size)
			if subStr != tt.want || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got %s, want %s and error is %+v", subStr, tt.want, err)
			}
		})
	}
}

func Test_GetTLVType(t *testing.T) {
	var tests = []struct {
		caseName string
		str      string
		tlvType  string
		err      error
	}{
		{
			"Success A",
			"A05",
			"A05",
			nil,
		},
		{
			"Success N",
			"N06",
			"N06",
			nil,
		},
		{
			"Error: require min 3 chars",
			"A0",
			"",
			errors.New("string require min 3 chars"),
		},
		{
			"Error: unknown type",
			"X05",
			"",
			errors.New("type invalid"),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s: %s", tt.caseName, tt.str)
		t.Run(testname, func(t *testing.T) {
			tlvType, err := getTLVType(tt.str)
			if tlvType != tt.tlvType || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got %s, want %s and error is %+v", tlvType, tt.tlvType, err)
			}
		})
	}
}

func Test_GetTLVLength(t *testing.T) {
	var tests = []struct {
		caseName  string
		str       string
		tlvLength string
		err       error
	}{
		{
			"Success",
			"11",
			"11",
			nil,
		},
		{
			"Error",
			"1",
			"",
			errors.New("string require min 2 chars"),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s: %s", tt.caseName, tt.str)
		t.Run(testname, func(t *testing.T) {
			tlvLength, err := getTLVLength(tt.str)
			if tlvLength != tt.tlvLength || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got %s, want %s and error is %+v", tlvLength, tt.tlvLength, err)
			}
		})
	}
}

func Test_ProcessTLVStr(t *testing.T) {
	var tests = []struct {
		caseName string
		strVal   string
		want     []map[string]string
		err      error
	}{
		{
			"Success",
			"A0511AB398765UJ1N230200",
			[]map[string]string{
				{"type": "A05", "length": "11", "value": "AB398765UJ1"},
				{"type": "N23", "length": "02", "value": "00"},
			},
			nil,
		},
		{
			"Error empty",
			"",
			nil,
			errors.New("str is empty"),
		},
		{
			"Error type",
			"X0511AB398765UJ1N230200",
			nil,
			errors.New("type invalid"),
		},
		{
			"Error type number",
			"ABC11AB398765UJ1N230200",
			nil,
			errors.New("type invalid number"),
		},
		{
			"Error length",
			"A051AAB398765UJ1N230200",
			nil,
			errors.New("length must be number"),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s: %s", tt.caseName, tt.strVal)
		t.Run(testname, func(t *testing.T) {
			tlvMap, err := processTLVStr(tt.strVal)
			if !reflect.DeepEqual(tlvMap, tt.want) || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got %+v, want %+v and error is %+v", tlvMap, tt.want, err)
			}
		})
	}
}
