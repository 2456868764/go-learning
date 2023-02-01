package container

import "testing"

func TestSet_Delete(t *testing.T) {
	type args struct {
		item SetItem
	}
	tests := []struct {
		name string
		s    Set
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSet_Has(t *testing.T) {
	type args struct {
		item SetItem
	}
	tests := []struct {
		name string
		s    Set
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Has(tt.args.item); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Insert(t *testing.T) {
	type args struct {
		item SetItem
	}
	tests := []struct {
		name string
		s    Set
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSet_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
