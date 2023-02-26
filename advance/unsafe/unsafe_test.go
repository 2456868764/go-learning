package unsafe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnSafeAccessor_FieldInt(t *testing.T) {

	tests := []struct {
		name    string
		entity  any
		field   string
		wantVal int
		wantErr error
	}{
		{
			name: "testNormal",
			entity: &TestData{D:9},
			field: "D",
			wantVal: 9,
			wantErr: nil,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUnsafeAccessor(tt.entity)
			assert.Nil(t, err)
			val, error:= u.FieldInt(tt.field)
			assert.Equal(t, tt.wantErr, error)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantVal, val)
		})
	}
}

func TestUnSafeAccessor_SetFieldInt(t *testing.T) {
	tests := []struct {
		name    string
		entity  *TestData
		field   string
		fieldVal int
		wantErr error
	}{
		{
			name: "testNormal",
			entity: &TestData{D:9},
			field: "D",
			fieldVal: 18,
			wantErr: nil,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUnsafeAccessor(tt.entity)
			assert.Nil(t, err)
			err=u.SetFieldInt(tt.field, tt.fieldVal)
			assert.Nil(t, err)
			assert.Equal(t, tt.fieldVal, tt.entity.D)
		})
	}
}

type TestData struct {
	D int
}
