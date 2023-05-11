package dao

import "testing"

func TestGetVideos(t *testing.T) {
	type args struct {
		time int64
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		tt.args.time = -1
		t.Run(tt.name, func(t *testing.T) {
			GetVideos(tt.args.time)
		})
	}
}
