package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	tests := []struct {
		name     string
		setData  []byte
		expected []byte
	}{
		{
			name:     "Set and Get data",
			setData:  []byte("test data"),
			expected: []byte("test data"),
		},
		{
			name:     "Set and Get empty data",
			setData:  []byte(""),
			expected: []byte(""),
		},
		{
			name:     "Set and Get nil data",
			setData:  nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := New()

			cache.Set(tt.setData)
			data, timestamp := cache.Get()

			require.Equal(t, tt.expected, data)
			require.NotEmpty(t, timestamp)
			require.Contains(t, timestamp, "GMT")
		})
	}
}
