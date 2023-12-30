package lecture8

import (
	"reflect"
	"testing"
)

func TestPaidStaircase(t *testing.T) {
	type args struct {
		n int
		p []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "complex case",
			args: args{
				n: 8,
				p: []int{0,3,2,4,6,1,1,5,3},
			},
			want: []int{0,2,3,5,6,8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaidStaircase(tt.args.n, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaidStaircase() = %v, want %v", got, tt.want)
			}
		})
	}
}
