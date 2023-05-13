package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loggestCommonPrefix(t *testing.T) {
	type testcase struct {
		in  []string
		out string
	}
	for _, tc := range []testcase{
		{
			in:  []string{},
			out: "",
		},
		{
			in:  []string{""},
			out: "",
		},
		{
			in:  []string{"", ""},
			out: "",
		},
		{
			in:  []string{"", "a"},
			out: "",
		},
		{
			in:  []string{"a", "", "b"},
			out: "",
		},
		{
			in:  []string{"a", "", "a"},
			out: "",
		},
		{
			in:  []string{"a", "", "aa"},
			out: "",
		},
		{
			in:  []string{"a", "a", "b"},
			out: "",
		},
		{
			in:  []string{"aa", "aa", "aab"},
			out: "aa",
		},
		{
			in:  []string{"flower", "flow", "flight"},
			out: "fl",
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.out, longestCommonPrefix(tc.in))
		})
	}
}
