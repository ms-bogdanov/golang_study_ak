package main

import (
	"reflect"
	"testing"
)

func TestCountWordsInText(t *testing.T) {
	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit
	amet ipsum mauris.
	Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor.
	Praesent et diam eget libero egestas mattis sit amet vitae augue.`

	type args struct {
		txt      string
		kyeWords []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "#1",
			args: args{txt: txt, kyeWords: []string{"sit", "amet", "lorem"}},
			want: map[string]int{"amet": 2, "lorem": 1, "sit": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountWordsInText(tt.args.txt, tt.args.kyeWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountWordsInText() = %v, want %v", got, tt.want)
			}
		})
	}
}
