package tencent

type TxOssStore struct {
}

func NewTxOssStore() *TxOssStore {
	return &TxOssStore{}
}

func (t *TxOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}