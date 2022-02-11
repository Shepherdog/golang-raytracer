package util

import (
	"raytracer/color"
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	cl := []color.Color{
		{10, 10, 10},
		{11, 11, 11},
		{12, 12, 12},
	}
	sumOfColor := func(accumulator color.Color, entry color.Color, idx int) color.Color {
		return color.Plus(accumulator, entry)
	}
	type args struct {
		source       interface{}
		initialValue interface{}
		reducer      interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"reduce color", args{source: cl, initialValue: color.Color{0, 0, 0}, reducer: sumOfColor}, color.Color{33, 33, 33}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reduce(tt.args.source, tt.args.initialValue, tt.args.reducer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reduce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
