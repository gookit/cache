package cache

import (
	"bytes"
	"encoding/gob"
)

// BindStruct get cache value and map to a struct
func BindStruct(val any, ptr any) error {
	// val must convert to byte
	return Unmarshal(val.([]byte), ptr)
}

// GobDecode decode data by gob.Decode
func GobDecode(bts []byte, ptr any) error {
	buf := bytes.NewBuffer(bts)
	dec := gob.NewDecoder(buf)

	return dec.Decode(ptr)
}

// GobEncode encode data by gob.Encode
func GobEncode(val any) (bs []byte, err error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err = enc.Encode(val); err != nil {
		return
	}

	return buf.Bytes(), nil
}
