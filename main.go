package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	// 程序内置
	endpint      = "oss-cn-chengdu.aliyuncs.com"
	accessKey    = "xx"
	accessSecret = "xx"

	// 默认配置
	bucketName = "devcloud-tool"

	// 用户需要传递的参数
	// 期望用户输入
	uploadFile = ""

	help = false
)

// upload 实现文件上传的函数
func upload(filePath string) error {
	// 1.实例化client
	client, err := oss.New(endpint, accessKey, accessSecret)
	if err != nil {
		return err
	}

	// 2.获取bucket对象
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 3.上传文件到该bucket
	return bucket.PutObjectFromFile(filePath, filePath)
}

// validate 参数合法性检查
func validate() error {
	if endpint == "" || accessKey == "" || accessSecret == "" {
		return errors.New("endpint, accessKey, accessSecret has one empty")
	}

	if uploadFile == "" {
		return errors.New("upload file path required")
	}

	return nil
}

func loadParams() {
	flag.BoolVar(&help, "h", false, "打印帮助信息")
	flag.StringVar(&uploadFile, "f", "", "上传文件的名称")
	flag.Parse()

	// 判断CLI 是否需要打印Help信息
	if help {
		usage()
		os.Exit(0)
	}
}

// 打印使用说明
func usage() {
	// 1. 打印一些描述信息
	fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <upload_file_path>
Options: 
`)
	// 2. 打印有哪些参数可以使用，就想-f
	flag.PrintDefaults()
}

func main() {
	// 参数加载
	loadParams()
	// 参数验证
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常: %s\n", err)
		usage()
		os.Exit(1)
	}
	if err := upload(uploadFile); err != nil {
		fmt.Printf("上传文件异常: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("文件: %s 上传完成\n", uploadFile)
}