package impl

import (
	"sync"
	"testing"
)

func TestLoadBalancer_GetPtr(t *testing.T) {
	type args struct {
		newTotal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should be ok 1",
			args: args{
				newTotal: 3,
			},
			want: 1,
		},
		{
			name: "should be ok 2",
			args: args{
				newTotal: 3,
			},
			want: 2,
		},
		{
			name: "should be ok 0",
			args: args{
				newTotal: 3,
			},
			want: 0,
		},
		{
			name: "should be ok 1",
			args: args{
				newTotal: 3,
			},
			want: 1,
		},
		{
			name: "should be ok 0",
			args: args{
				newTotal: 2,
			},
			want: 0,
		},
		{
			name: "should be ok 1",
			args: args{
				newTotal: 2,
			},
			want: 1,
		},
		{
			name: "should be ok 0",
			args: args{
				newTotal: 1,
			},
			want: 0,
		},
		{
			name: "should be ok 1",
			args: args{
				newTotal: 2,
			},
			want: 1,
		},
		{
			name: "should be ok 0",
			args: args{
				newTotal: 2,
			},
			want: 0,
		},
	}

	l := &LoadBalancer{
		Mtx: &sync.Mutex{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := l.GetNextPtr(tt.args.newTotal); got != tt.want {
				t.Errorf("LoadBalancer.GetPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
