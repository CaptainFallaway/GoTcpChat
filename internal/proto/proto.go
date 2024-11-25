package proto

import (
	"bytes"
	"encoding/gob"
	"errors"
)

var (
	ErrInvalidHeader   = errors.New("invalid header bytes, must be at least 2 bytes long")
	ErrVersionMismatch = errors.New("version mismatch")
)

func GetOpCode(bytes []byte, expectedVersion uint8) (uint8, error) {
	if len(bytes) < 2 {
		return 0, ErrInvalidHeader
	}

	version, op := bytes[0], bytes[1]

	if version != expectedVersion {
		return 0, ErrVersionMismatch
	}

	return op, nil
}

func Decode[T any](buf []byte) (*T, error) {
	reader := bytes.NewBuffer(buf[2:])

	decoder := gob.NewDecoder(reader)

	var t *T = new(T)
	err := decoder.Decode(t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func Encode[T any](msg *T, op uint8, version uint8) ([]byte, error) {
	var msgBuf bytes.Buffer
	encoder := gob.NewEncoder(&msgBuf)

	err := encoder.Encode(msg)
	if err != nil {
		return nil, err
	}

	retBuf := []byte{version, op}
	retBuf = append(retBuf, msgBuf.Bytes()...)

	return retBuf, nil
}
