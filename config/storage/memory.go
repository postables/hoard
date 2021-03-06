package storage

func NewMemoryConfig(addressEncoding string) *StorageConfig {
	return NewStorageConfig(Memory, addressEncoding)
}

func DefaultMemoryConfig() *StorageConfig {
	return NewMemoryConfig(DefaultAddressEncodingName)
}
