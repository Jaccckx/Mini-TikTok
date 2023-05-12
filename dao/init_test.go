package dao

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {
	Init()
	logrus.Debugln("InitDeps successfully!")
}
