package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateList(t *testing.T) {
	type args struct {
		object any
	}
	tests := []struct {
		name    string
		args    args
		want    []any
		wantErr error
	}{
		// TODO: Add test cases.
		{
			"nil",
			args{nil},
			nil,
			errors.New("type is not supported"),
		},
		{
			"array",
			args{[3]int{1, 2, 3}},
			[]any{1, 2, 3},
			nil,
		},
		{
			"arrayString",
			args{[3]string{"ab", "cd", "ef"}},
			[]any{"ab", "cd", "ef"},
			nil,
		},
		{
			"slice",
			args{[]int{1, 2, 3, 4}},
			[]any{1, 2, 3, 4},
			nil,
		},
		//{
		//	"string",
		//	args{"hello"},
		//	[]any{'h', 'e', 'l', 'l', 'o'},
		//	nil,
		//},
		//{
		//	"stringUTF8",
		//	args{"结构体之间"},
		//	[]any{"结", "构", "体", "之", "间"},
		//	nil,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateList(tt.args.object)

			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tt.want, got)

		})
	}
}

func TestIterateMap(t *testing.T) {
	type args struct {
		object any
	}
	tests := []struct {
		name    string
		args    args
		wantKeys    []any
		wantValues   []any
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:    "nil",
			args:  args{nil},
			wantErr: errors.New("type is not supported"),
		},
		{
			name: "good",
			args: args{map[string]string{
				"a_k_1": "a_v_1",
			}},
			wantKeys:   []any{"a_k_1"},
			wantValues: []any{"a_v_1"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, values, err := IterateMap(tt.args.object)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantKeys, keys)
			assert.Equal(t, tt.wantValues, values)
		})
	}
}
