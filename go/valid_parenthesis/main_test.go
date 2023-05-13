package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type testcase struct {
		string
		bool
	}
	for _, tc := range []testcase{
		{"[](){}", true},
		{"[]{}", true},
		{"[]{}", true},
		{"[]{}(", false},
		{"[]{})", false},
		{"{(})", false},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.bool, isValid(tc.string))
		})
	}
}
