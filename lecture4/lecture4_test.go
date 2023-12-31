package lecture4

import "testing"

func Test_climbStiars(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test 1",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "simple test 1",
			args: args{n: 5},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStiars() = %v, want %v", got, tt.want)
			}
		})
	}
}
