package cli

import (
	"errors"
	"github.com/ives22/cloud_station/store"
	"github.com/ives22/cloud_station/store/aliyun"
	"github.com/ives22/cloud_station/store/aws"
	"github.com/ives22/cloud_station/store/tencent"
	"github.com/spf13/cobra"
)

var (
	ossProvier   string
	ossEndPoint  string
	accessKey    string
	accessSecret string
	bucketName   string
	uploadFile   string
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件上传",
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)
		switch ossProvier {
		case "aliyun":
			uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
				EndPoint:     ossEndPoint,
				AccessKey:    accessKey,
				AccessSecret: accessSecret,
			})
		case "tx":
			uploader = tencent.NewTxOssStore()
		case "aws":
			uploader = aws.NewAwsOssStore()
		default:
			return errors.New("not support oss storage provider")
		}
		if err != nil {
			return err
		}

		// 使用uploader 来上传文件
		return uploader.Upload(bucketName, uploadFile, uploadFile)
	},
}

func init() {
	// 把upload命令作为root命令的子命令
	UploadCmd.PersistentFlags().StringVarP(&ossProvier, "provider", "p", "aliyun", "oss storage provier [aliyun/tencent/aws]")
	UploadCmd.PersistentFlags().StringVarP(&ossEndPoint, "endpoint", "e", "oss-cn-chengdu.aliyuncs.com", "oss storage provier endpoint")
	UploadCmd.PersistentFlags().StringVarP(&bucketName, "bucket_name", "b", "devcloud-tool", "oss storage provier bucket name")
	UploadCmd.PersistentFlags().StringVarP(&accessKey, "access_key", "k", "", "oss storage provier ak")
	UploadCmd.PersistentFlags().StringVarP(&accessSecret, "access_secret", "s", "", "oss storage provier sk")
	UploadCmd.PersistentFlags().StringVarP(&uploadFile, "upload_file", "f", "", "upload file name")
	RootCmd.AddCommand(UploadCmd)
}