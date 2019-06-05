package strexp

import (
	"testing"
)

func Test_numFromStart(t *testing.T) {

	tests := []struct {
		name  string
		arg   string
		want  int
		want1 string
	}{
		{"empty", "", 1, ""},
		{"abc", "abc", 1, "abc"},
		{"33abc", "33abc", 33, "abc"},
		{"345", "345", 345, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := numFromStart(tt.arg)
			if got != tt.want {
				t.Errorf("numFromStart() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("numFromStart() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStrExpand(t *testing.T) {
	tests := []struct {
		arg string
		want string
	}{
		{"", ""},
		{"45", ""},
		{"a", "a"},
		{"a3", "aaa"},
		{"a12", "aaaaaaaaaaaa"},
		{"ba3", "baaa"},
		{"abcd", "abcd"},
		{"a4bc2d5e", "aaaabccddddde"},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := StrExpand(tt.arg); got != tt.want {
				t.Errorf("StrExpand() = %v, want '%v'", got, tt.want)
			}
		})
	}
}
