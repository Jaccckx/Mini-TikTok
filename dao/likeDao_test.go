package dao

import (
	"testing"
)

func TestIsFavoriteVideo(t *testing.T) {
	Init()
	type args struct {
		userID  int64
		videoID int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsFavoriteVideo",
			args: args{userID: 44, videoID: 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IsFavoriteVideo(tt.args.userID, tt.args.videoID); got != tt.want {
				t.Errorf("IsFavoriteVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActionLikeVideo(t *testing.T) {
	Init()
	type args struct {
		userID     int64
		videoID    int64
		actionType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestActionLikeVideo",
			args:    args{userID: 44, videoID: 1, actionType: "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ActionLikeVideo(tt.args.userID, tt.args.videoID, tt.args.actionType); (err != nil) != tt.wantErr {
				t.Errorf("ActionLikeVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFavoriteCountByUserID(t *testing.T) {
	Init()
	type args struct {
		userID int64
	}
	tests := []struct {
		name      string
		args      args
		wantCount int64
	}{
		{
			name:      "TestGetFavoriteCountByUserID",
			args:      args{userID: 50},
			wantCount: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount, _ := GetFavoriteCountByUserID(tt.args.userID); gotCount != tt.wantCount {
				t.Errorf("GetFavoriteCountByUserID() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestGetFavoriteCountByVideoID(t *testing.T) {
	Init()
	type args struct {
		videoID int64
	}
	tests := []struct {
		name      string
		args      args
		wantCount int64
		wantErr   bool
	}{
		{
			name:      "TestGetFavoriteCount",
			args:      args{videoID: 1},
			wantCount: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, err := GetFavoriteCountByVideoID(tt.args.videoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFavoriteCountByVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("GetFavoriteCountByVideoID() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
