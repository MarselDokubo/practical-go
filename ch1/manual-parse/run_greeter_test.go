package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunGreeter(t *testing.T) {
	cases := []struct {
		config         greeterConfig
		input          string
		expectedOutput string
		expectedErr    error
	}{
		{
			config:         greeterConfig{printUsage: true},
			expectedOutput: greetDesc,
		},
		{
			config:         greeterConfig{numTimes: 1},
			input:          "",
			expectedOutput: strings.Repeat("Welcome. Enter your name: ", 1),
			expectedErr:    ErrNoName,
		},
		{
			config:         greeterConfig{numTimes: 1},
			input:          "Marsel Dokubo",
			expectedOutput: "Welcome. Enter your name: \n" + "Howdy Marsel Dokubo\n",
		},
	}

	for _, tc := range cases {
		r := strings.NewReader(tc.input)
		outBuf := new(bytes.Buffer)
		err := runGreeter(r, outBuf, tc.config)
		if tc.expectedErr == nil && err != nil {
			t.Fatalf("Expected nil error, got: %v", err)
		}
		if tc.expectedErr != nil && tc.expectedErr.Error() != err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v", tc.expectedErr.Error(), err.Error())
		}
		if tc.expectedErr != nil && err == nil {
			t.Fatalf("Expected error to be: %v, recieved nil error", tc.expectedErr)
		}

		op := outBuf.String()

		if op != tc.expectedOutput {
			t.Errorf("Expected stdout message to be: %v, got: %v\n", tc.expectedOutput, op)
		}
		outBuf.Reset()
	}

}
