package main_test

import (
	"testing"

	. "github.com/Manuelshub/RSA"
	"github.com/stretchr/testify/assert"
)

func TestGetFileContent(t *testing.T) {
	var tests = []struct {
		file     string
		expected string
	}{
		{"test.txt", "Hello, World!"},
		{"test2.txt", "This is just a test file!"},
		{"test3.txt(", ""},
	}
	for _, test := range tests {
		if got := GetFileContent(test.file); got != test.expected {
			t.Errorf("Expected: %v, Got: %v\n", test.expected, got)
		}
	}
	assert.Equal(t, GetFileContent("test.txt"), "Hello, World!")
	assert.Equal(t, GetFileContent("test2.txt"), "This is just a test file!")
	assert.NotEqual(t, GetFileContent("test3.txt"), "Hello, World!")
}

func TestGetFactors(t *testing.T) {
	var tests = []struct {
		n        uint64
		expected [2]uint64
	}{
		{4, [2]uint64{2, 2}},
		{12, [2]uint64{2, 6}},
		{1024, [2]uint64{2, 512}},
		{4958, [2]uint64{2, 2479}},
		// {1718944270642558716715, [2]int{5, 343788854128511743343}},
	}
	for _, test := range tests {
		if got := GetFactors(test.n); got != test.expected {
			t.Errorf("Expected: %v, Got: %v\n", test.expected, got)
		} else {
			t.Logf("Expected: %v, Got: %v\n", test.expected, got)
		}
	}
	assert.Equal(t, GetFactors(4), [2]int{2, 2})
	assert.Equal(t, GetFactors(12), [2]int{2, 6})
	assert.NotEqual(t, GetFactors(1024), [2]int{2, 514})
}
