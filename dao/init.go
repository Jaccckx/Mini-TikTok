package dao

import (
	"fmt"
	"mini-tiktok/config"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var ossBucket *oss.Bucket
var urlPrefix string

func DataBaseInit() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/tiktok?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBUrl)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}

func OssInit() {

	// 设置Endpoint和BucketName
	endpoint := "oss-cn-beijing.aliyuncs.com"
	bucketName := "mini-tiktok-bytedance"
	urlPrefix = "https://" + bucketName + "." + endpoint + "/"
	// 创建OSS客户端
	client, err := oss.New(endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 获取Bucket
	ossBucket, err = client.Bucket(bucketName)
	_ = ossBucket
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func Init() {
	DataBaseInit()
	OssInit()
}
