package tabletest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Ilham)",
			request:  "Ilham",
			expected: "Hello Ilham",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
