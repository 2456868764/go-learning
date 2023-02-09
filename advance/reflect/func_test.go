package reflect

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateMethods(t *testing.T) {

	tests := []struct {
		name    string
		object  any
		want    map[string]*MethodInfo
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:   "VALID: normal struct case",
			object: Blog{},
			want: map[string]*MethodInfo{
				"GetId": {
					Name:      "GetId",
					ParamsIn:  []reflect.Type{reflect.TypeOf(Blog{})},
					ParamsOut: []reflect.Type{reflect.TypeOf(int64(0))},
				},
			},
		},
		{
			name:   "VALID: pointer case",
			object: &Blog{},
			want: map[string]*MethodInfo{
				"GetId": {
					Name:      "GetId",
					ParamsIn:  []reflect.Type{reflect.TypeOf(&Blog{})},
					ParamsOut: []reflect.Type{reflect.TypeOf(int64(0))},
				},
				"ChangeTitle": {
					Name:      "ChangeTitle",
					ParamsIn:  []reflect.Type{reflect.TypeOf(&Blog{}), reflect.TypeOf("")},
					ParamsOut: []reflect.Type{},
				},
				"IncreaseThumbUp": {
					Name:      "IncreaseThumbUp",
					ParamsIn:  []reflect.Type{reflect.TypeOf(&Blog{})},
					ParamsOut: []reflect.Type{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateMethods(tt.object)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			//fmt.Printf("w %+v", tt.want["ChangeTitle"])
			//fmt.Printf("w %+v", got["ChangeTitle"])
			assert.Equal(t, tt.want, got)
		})
	}
}
