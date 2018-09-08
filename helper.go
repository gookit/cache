package cache

// GetStruct get cache value and map to a struct
func GetStruct(val interface{}, ptr interface{}) error {
	// val must convert to byte
	return Unmarshal(val.([]byte), ptr)
}
