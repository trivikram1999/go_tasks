package operations

import (
	"math"
	"testing"
)

func TestPerformOperation(t *testing.T) {
	type args struct {
		op1 float64
		op2 float64
		opt string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "add_test",
			args: args{
				op1: 10,
				op2: 22,
				opt: "+",
			},
			want:    32,
			wantErr: false,
		},
		{
			name: "sub_test",
			args: args{
				op1: 100,
				op2: 20,
				opt: "-",
			},
			want:    80,
			wantErr: false,
		},
		{
			name: "mul_test",
			args: args{
				op1: 10,
				op2: 2,
				opt: "*",
			},
			want:    20,
			wantErr: false,
		},
		{
			name: "div_test",
			args: args{
				op1: 100,
				op2: 2,
				opt: "/",
			},
			want:    50,
			wantErr: false,
		},
		{
			name: "div0_test",
			args: args{
				op1: 10,
				op2: 0,
				opt: "/",
			},
			want:    math.Inf(+1),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PerformOperation(tt.args.op1, tt.args.op2, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("PerformOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PerformOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
