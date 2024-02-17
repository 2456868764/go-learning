package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {

	tests := []struct {
		name string
		arg  []byte
		want []byte
	}{
		{
			name: "test1",
			arg:  []byte("hello"),
			want: []byte("olleh"),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.arg)
			assert.Equal(t, tt.want, tt.arg)
		})
	}
}
