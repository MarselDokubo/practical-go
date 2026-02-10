package main

import (
	"errors"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	cases := []struct {
		config      greeterConfig
		expectedErr error
	}{
		{
			config:      greeterConfig{},
			expectedErr: errors.New("Invalid argument. Must specify a number greater than 0."),
		},
		{
			config:      greeterConfig{numTimes: -1},
			expectedErr: errors.New("Invalid argument. Must specify a number greater than 0."),
		},
		{
			config:      greeterConfig{numTimes: 1},
			expectedErr: nil,
		},
	}

	for _, tc := range cases {
		err := validateArgs(tc.config)
		if tc.expectedErr != nil && tc.expectedErr.Error() != err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.expectedErr, err)
		}
		if tc.expectedErr == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
	}
}
