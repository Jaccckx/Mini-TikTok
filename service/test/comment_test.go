package test

import (
	"mini-tiktok/dao"
	"mini-tiktok/service"
	"testing"
)

func TestComment(t *testing.T) {
	csi := service.CommentServiceImpl{}

	id, err := csi.CreateComment(1, 1, "test comment")
	if err != nil {
		t.Fatal(err)
	}

	err = csi.DeleteCommentByID(id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	dao.Init()

	m.Run()
}
