package lecture7

import "testing"

func TestPaidStaircase(t *testing.T) {
	type args struct {
		n int
		p []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case #1",
			args: args{
				n: 3,
				p: []int{0,3,2,4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaidStaircase(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("PaidStaircase() = %v, want %v", got, tt.want)
			}
		})
	}
}
