package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_countAll(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want Output
	}{
		{name: "Happy Path",
			args: args{r: strings.NewReader("Character Counter is a 100% free online character count calculator that's simple to use\nTest line")},
			want: Output{CharacterCount: 96, WordCount: 16, LineCount: 2},
		},
		{name: "Empty Input",
			args: args{r: strings.NewReader("")},
			want: Output{CharacterCount: 0, WordCount: 0, LineCount: 0},
		},
		{name: "Single Line with Spaces",
			args: args{r: strings.NewReader("hello   world")},
			want: Output{CharacterCount: 13, WordCount: 2, LineCount: 1},
		},
		{name: "Multiple Newlines",
			args: args{r: strings.NewReader("line1\n\nline3")},
			want: Output{CharacterCount: 10, WordCount: 2, LineCount: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countAll(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
