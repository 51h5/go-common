package utils

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name        string
		a           int
		b           int
		expectedVal int
	}{
		{
			name:        "both positive number and a < b",
			a:           1,
			b:           99,
			expectedVal: 99,
		},
		{
			name:        "both positive number and a > b",
			a:           99,
			b:           1,
			expectedVal: 99,
		},
		{
			name:        "both positive number and a = b",
			a:           99,
			b:           99,
			expectedVal: 99,
		},
		{
			name:        "both negative number and a < b",
			a:           -99,
			b:           -1,
			expectedVal: -1,
		},
		{
			name:        "both negative number and a > b",
			a:           -1,
			b:           -99,
			expectedVal: -1,
		},
		{
			name:        "negative or positive number and a < b",
			a:           -99,
			b:           99,
			expectedVal: 99,
		},
		{
			name:        "negative or positive number and a > b",
			a:           99,
			b:           -99,
			expectedVal: 99,
		},
		{
			name:        "positive or max number and b is max",
			a:           99,
			b:           math.MaxInt,
			expectedVal: math.MaxInt,
		},
		{
			name:        "positive or max number and a is max",
			a:           math.MaxInt,
			b:           99,
			expectedVal: math.MaxInt,
		},
		{
			name:        "negative or max number and b is max",
			a:           -99,
			b:           math.MaxInt,
			expectedVal: math.MaxInt,
		},
		{
			name:        "negative or max number and a is max",
			a:           math.MaxInt,
			b:           -99,
			expectedVal: math.MaxInt,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedVal, Max(tc.a, tc.b))
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name        string
		a           int
		b           int
		expectedVal int
	}{
		{
			name:        "both positive number and a < b",
			a:           1,
			b:           99,
			expectedVal: 1,
		},
		{
			name:        "both positive number and a > b",
			a:           99,
			b:           1,
			expectedVal: 1,
		},
		{
			name:        "both positive number and a = b",
			a:           99,
			b:           99,
			expectedVal: 99,
		},
		{
			name:        "both negative number and a < b",
			a:           -99,
			b:           -1,
			expectedVal: -99,
		},
		{
			name:        "both negative number and a > b",
			a:           -1,
			b:           -99,
			expectedVal: -99,
		},
		{
			name:        "negative or positive number and a < b",
			a:           -99,
			b:           99,
			expectedVal: -99,
		},
		{
			name:        "negative or positive number and a > b",
			a:           99,
			b:           -99,
			expectedVal: -99,
		},
		{
			name:        "positive or min number and b is min",
			a:           99,
			b:           math.MinInt,
			expectedVal: math.MinInt,
		},
		{
			name:        "positive or min number and a is min",
			a:           math.MinInt,
			b:           99,
			expectedVal: math.MinInt,
		},
		{
			name:        "negative or min number and b is min",
			a:           -99,
			b:           math.MinInt,
			expectedVal: math.MinInt,
		},
		{
			name:        "negative or min number and a is min",
			a:           math.MinInt,
			b:           -99,
			expectedVal: math.MinInt,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedVal, Min(tc.a, tc.b))
		})
	}
}

func TestMax64(t *testing.T) {
	testCases := []struct {
		name        string
		a           int64
		b           int64
		expectedVal int64
	}{
		{
			name:        "both positive number and a < b",
			a:           1,
			b:           99,
			expectedVal: 99,
		},
		{
			name:        "both positive number and a > b",
			a:           99,
			b:           1,
			expectedVal: 99,
		},
		{
			name:        "both positive number and a = b",
			a:           99,
			b:           99,
			expectedVal: 99,
		},
		{
			name:        "both negative number and a < b",
			a:           -99,
			b:           -1,
			expectedVal: -1,
		},
		{
			name:        "both negative number and a > b",
			a:           -1,
			b:           -99,
			expectedVal: -1,
		},
		{
			name:        "negative or positive number and a < b",
			a:           -99,
			b:           99,
			expectedVal: 99,
		},
		{
			name:        "negative or positive number and a > b",
			a:           99,
			b:           -99,
			expectedVal: 99,
		},
		{
			name:        "positive or max number and b is max",
			a:           99,
			b:           math.MaxInt64,
			expectedVal: math.MaxInt64,
		},
		{
			name:        "positive or max number and a is max",
			a:           math.MaxInt64,
			b:           99,
			expectedVal: math.MaxInt64,
		},
		{
			name:        "negative or max number and b is max",
			a:           -99,
			b:           math.MaxInt64,
			expectedVal: math.MaxInt64,
		},
		{
			name:        "negative or max number and a is max",
			a:           math.MaxInt64,
			b:           -99,
			expectedVal: math.MaxInt64,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedVal, Max64(tc.a, tc.b))
		})
	}
}

func TestMin64(t *testing.T) {
	testCases := []struct {
		name        string
		a           int64
		b           int64
		expectedVal int64
	}{
		{
			name:        "both positive number and a < b",
			a:           1,
			b:           99,
			expectedVal: 1,
		},
		{
			name:        "both positive number and a > b",
			a:           99,
			b:           1,
			expectedVal: 1,
		},
		{
			name:        "both positive number and a = b",
			a:           99,
			b:           99,
			expectedVal: 99,
		},
		{
			name:        "both negative number and a < b",
			a:           -99,
			b:           -1,
			expectedVal: -99,
		},
		{
			name:        "both negative number and a > b",
			a:           -1,
			b:           -99,
			expectedVal: -99,
		},
		{
			name:        "negative or positive number and a < b",
			a:           -99,
			b:           99,
			expectedVal: -99,
		},
		{
			name:        "negative or positive number and a > b",
			a:           99,
			b:           -99,
			expectedVal: -99,
		},
		{
			name:        "positive or min number and b is min",
			a:           99,
			b:           math.MinInt64,
			expectedVal: math.MinInt64,
		},
		{
			name:        "positive or min number and a is min",
			a:           math.MinInt64,
			b:           99,
			expectedVal: math.MinInt64,
		},
		{
			name:        "negative or min number and b is min",
			a:           -99,
			b:           math.MinInt64,
			expectedVal: math.MinInt64,
		},
		{
			name:        "negative or min number and a is min",
			a:           math.MinInt64,
			b:           -99,
			expectedVal: math.MinInt64,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedVal, Min64(tc.a, tc.b))
		})
	}
}
