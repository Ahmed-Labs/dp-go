package lecture5

import (
	"testing"
)

func Test_climbStairs3Steps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{n: 3},
			want: 4,
		},
		{
			name: "test 2",
			args: args{n: 5},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs3Steps(tt.args.n); got != tt.want {
				t.Errorf("climbStairs3Steps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs3StepsOpt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{n: 3},
			want: 4,
		},
		{
			name: "test 2",
			args: args{n: 5},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs3StepsOpt(tt.args.n); got != tt.want {
				t.Errorf("climbStairs3StepsOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}
