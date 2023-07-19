package aliyun_test

import (
	"github.com/ives22/cloud_station/store"
	"github.com/ives22/cloud_station/store/aliyun"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	uploader     store.Uploader
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

// TestUpload Aliyun Oss Store Upload 测试用例
func TestUpload(t *testing.T) {
	// 使用 assert 编写测试用例的断言语句
	// 通过New获取一个断言实例
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		// 没有Error 开始下一个步骤
		t.Log("upload ok")
	}
}

func TestUploadError(t *testing.T) {
	// 通过New获取一个断言实例
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "bbb.go")
	if should.Error(err, "no such file or directory") {
		t.Log("upload ok")
	}

}

// 通过init 编写uploader 实例化逻辑
func init() {
	ali, err := aliyun.NewDefaultAliOssStore()
	if err != nil {
		panic(err)
	}
	uploader = ali
}