package backup

type HDFS struct {
	Address string
}

type StorageBackend_HDFS struct {
	HDFS *HDFS
	Local *Local
}

func (*StorageBackend_HDFS) isStorageBackend_Backend()         {}

func (*StorageBackend_HDFS) MarshalTo([]byte) (int, error) {
	return 0,nil
}
func (*StorageBackend_HDFS) Size() int {
	return 0
}