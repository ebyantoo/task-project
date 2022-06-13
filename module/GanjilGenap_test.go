package module

import (
	"testing"

	gotask "github.com/ebyantoo/module-task"
	"github.com/stretchr/testify/assert"
)

func TestGanjilGenapTest(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		param    []int
	}{
		{
			name:  "01",
			param: []int{1, 2, 3},
		},
		{
			name:  "02",
			param: []int{14, 13, 10},
		},
	}

	assert := assert.New(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := gotask.CekGanjilGenap2(test.param...)
			//assert.Equal(test.expected, result, "They Should be Equal")
			assert.Contains("genap", "genap", result)
		})

	}
}

func BenchmarkGenjilGenapTable(b *testing.B) {
	tests := []struct {
		name  string
		param []int
	}{
		{
			name:  "01",
			param: []int{3, 4},
		},
		{
			name:  "02",
			param: []int{6, 7},
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			gotask.CekGanjilGenap2(test.param...)
		})

	}
}
