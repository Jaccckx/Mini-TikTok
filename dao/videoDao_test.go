package dao

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetVideoCount(t *testing.T) {
	type args struct {
		userID int64
	}
	tests := []struct {
		name      string
		args      args
		wantCount int64
		wantErr   bool
	}{
		{
			name:      "TestGetVideoCount01",
			args:      args{userID: 1},
			wantCount: 2,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, err := GetVideoCount(tt.args.userID)
			if (err != nil) != tt.wantErr {
				logrus.Printf("GetVideoCount()() = %v, want %v", gotCount, tt.wantCount)
				t.Errorf("GetVideoCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				logrus.Printf("name:%v,GetVideoCount() = %v, want %v", tt.name, gotCount, tt.wantCount)
				t.Errorf("GetVideoCount() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestUploadFileToOss(t *testing.T) {
	Init()
	type args struct {
		path  string
		title string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "UploadFileToOss",
			args:    args{path: "../resources/upload/video4Test.mp4", title: "Test.mp4"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UploadFileToOss(tt.args.path, tt.args.title); (err != nil) != tt.wantErr {
				t.Errorf("UploadFileToOss() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteFileFromOss(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "DeleteFileFromOss",
			args:    args{title: "Game.mp4"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFileFromOss(tt.args.title); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFileFromOss() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUrlFromOss(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name        string
		args        args
		wantPlayURL string
		wantErr     bool
	}{
		{
			name:        "TestGetUrlFromOss",
			args:        args{fileName: "The Long Season.mp4"},
			wantPlayURL: "http://mini-tiktok-bytedance.oss-cn-beijing.aliyuncs.com/The%20Long%20Season.mp4",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlayURL, err := GetUrlFromOss(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUrlFromOss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPlayURL[:len(tt.wantPlayURL)] != tt.wantPlayURL {
				t.Errorf("GetUrlFromOss() = %v, want %v", gotPlayURL[:len(tt.wantPlayURL)], tt.wantPlayURL)
			}
		})
	}
}

func TestInsertVideoRecordToDataBase(t *testing.T) {
	type args struct {
		title    string
		userID   int64
		playUrl  string
		coverURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestInsertVideoRecordToDataBase",
			args: args{title: "test", userID: 1, playUrl: "https://mini-tiktok-bytedance.oss-cn-beijing.aliyuncs.com/manchester%20by%20the%20sea.mp4",
				coverURL: "https://mini-tiktok-bytedance.oss-cn-beijing.aliyuncs.com/manchester%20by%20the%20sea.mp4"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var videoID int64
			var err error
			if videoID, err = InsertVideoRecordToDataBase(tt.args.title, tt.args.userID, tt.args.playUrl, tt.args.coverURL); (err != nil) != tt.wantErr {
				t.Errorf("InsertVideoRecordToDataBase() error = %v, wantErr %v", err, tt.wantErr)
			}
			DeleteVideoRecordFromDataBase(videoID)
		})
	}
}

func TestGetVideoListByUserID(t *testing.T) {
	Init()
	type args struct {
		userID int64
	}
	tests := []struct {
		name            string
		args            args
		wantTableVideos []TableVideo
	}{
		{
			name: "TestGetVideoListByUserID",
			args: args{userID: 44},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTableVideos := GetVideoListByUserID(tt.args.userID); len(gotTableVideos) != 1 {
				t.Errorf("GetVideoListByUserID() = %v, want %v", gotTableVideos, tt.wantTableVideos)
			}
		})
	}
}
