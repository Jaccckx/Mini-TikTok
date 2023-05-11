package service

import (
	"log"
	"math"
	"mini-tiktok/dao"
	"testing"
	"time"
)

func TestVideoServiceImpl_GetVideos(t *testing.T) {
	v := VideoServiceImpl{}
	dao.Init()
	var videos []Video
	var NextTime int64
	if videos, NextTime = v.GetVideos(time.Now(), -1); len(videos) == 0 || NextTime == 0 {
		t.Errorf("wrong!\n")
	} else {
		log.Println("videos-num: ", len(videos))
		log.Printf("%+v\n", videos)
		log.Println("NextTime: ", NextTime)
	}

	if videos, NextTime = v.GetVideos(time.Unix(NextTime, 0), -1); len(videos) != 0 || NextTime != math.MaxInt64 {
		t.Errorf("wrong!\n")
	} else {
		log.Println("yes empty videos!, NextTime == MaxInt64")
	}

}
