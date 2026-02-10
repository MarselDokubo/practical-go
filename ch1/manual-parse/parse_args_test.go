package main

import (
	"errors"
	"testing"
)

type testConfig struct {
	args         []string
	expectedConf greeterConfig
	expectedErr  error
}

func TestParseArgs(t *testing.T) {
	cases := []testConfig{
		{
			args:         []string{"-h"},
			expectedConf: greeterConfig{printUsage: true, numTimes: 0},
			expectedErr:  nil,
		},
		{
			args:         []string{"10"},
			expectedConf: greeterConfig{printUsage: false, numTimes: 10},
			expectedErr:  nil,
		},
		{
			args:         []string{"ab"},
			expectedConf: greeterConfig{printUsage: false, numTimes: 0},
			expectedErr:  errors.New("strconv.Atoi: parsing \"ab\": invalid syntax"),
		},
		{
			args:         []string{"10", "foo"},
			expectedConf: greeterConfig{printUsage: false, numTimes: 0},
			expectedErr:  errors.New("Invalid number of arguments."),
		},
	}

	for _, tc := range cases {
		c, err := parseArgs(tc.args)
		if tc.expectedErr != nil && tc.expectedErr.Error() != err.Error() {
			t.Fatalf("Expected error to be: %v, got %v\n", tc.expectedErr, err)
		}
		if tc.expectedErr == nil && err != nil {
			t.Fatalf("Expected error to be: %v, got %v\n", tc.expectedErr, err)
		}
		if tc.expectedConf.printUsage != c.printUsage {
			t.Fatalf("Expected printUsage to be: %v, got %v\n", tc.expectedConf.printUsage, c.printUsage)
		}
		if tc.expectedConf.numTimes != c.numTimes {
			t.Fatalf("Expected numTimes to be: %v, got %v\n", tc.expectedConf.numTimes, c.numTimes)
		}
	}
}
