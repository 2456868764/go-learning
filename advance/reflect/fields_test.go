package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateFields(t *testing.T) {

	tests := []struct {
		name    string
		object  any
		want    map[string]any
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "VALID:normal case",
			object: Blog{
				Id:      12,
				Title:   "demo blog",
				ThumbUp: 1,
			},
			want: map[string]any{
				"Id":      int64(12),
				"Title":   "demo blog",
				"ThumbUp": int64(1),
				"author":  "",
			},
		},

		{
			name: "VALID:pointer case",
			object: &Blog{
				Id:      12,
				Title:   "demo blog",
				ThumbUp: 1,
			},
			want: map[string]any{
				"Id":      int64(12),
				"Title":   "demo blog",
				"ThumbUp": int64(1),
				"author":  "",
			},
		},
		{
			// invalid input
			name:    "ERROR:slice case",
			object:  []string{},
			wantErr: errors.New("type is not supported"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFields(tt.object)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)

		})
	}
}
