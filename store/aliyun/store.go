package aliyun

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ives22/cloud_station/store"
	"os"
)

var (
	// 对象是否实现了接口的约束
	_ store.Uploader = &AliOssStore{}
)

type AliOssStore struct {
	client *oss.Client
}

type Options struct {
	EndPoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) Validate() error {
	if o.EndPoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return errors.New("endPoint, accessKey, accessSecret has one empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		EndPoint:     os.Getenv("ALI_OSS_ENDPOINT"),
		AccessKey:    os.Getenv("ALI_AK"),
		AccessSecret: os.Getenv("ALI_SK"),
	})
}

// NewAliOssStore AliOssStore对象的构造函数
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	// 校验参数
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	c, err := oss.New(opts.EndPoint, opts.AccessKey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client: c,
	}, nil
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	// 2.获取bucket对象
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 3.上传文件到该bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return err
	}

	// 4.打印下载链接
	downloadUrl, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL: %s\n", downloadUrl)
	fmt.Println("请在1天之内下载.")
	return nil
}