package aws

type AwsOssStore struct {
}

func NewAwsOssStore() *AwsOssStore {
	return &AwsOssStore{}
}
func (a *AwsOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}