package helper

import (
	"fmt"
	"testing"
)

func TestProgressor(t *testing.T) {
	type args struct {
		ch chan uint8
	}

	chProgress := make(chan uint8)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "should be ok",
			args: args{
				ch: chProgress,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go Progressor(tt.args.ch)

			for v := range tt.args.ch {
				fmt.Print(v)
			}

		})
	}
}
