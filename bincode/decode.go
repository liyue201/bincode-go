package bincode

import (
	"encoding/binary"
	"io"
)

func DecodeBool(reader io.Reader) (bool, error) {
	var p [1]byte
	_, err := reader.Read(p[:])
	if err != nil {
		return false, err
	}
	return p[0] == 1, nil
}

func DecodeU8(reader io.Reader) (uint8, error) {
	var p [1]byte
	_, err := reader.Read(p[:])
	if err != nil {
		return 0, err
	}
	return p[0], nil
}

func DecodeU16(reader io.Reader) (uint16, error) {
	var p [2]byte
	_, err := reader.Read(p[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(p[:]), nil
}

func DecodeU32(reader io.Reader) (uint32, error) {
	var p [4]byte
	_, err := reader.Read(p[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(p[:]), nil
}

func DecodeU64(reader io.Reader) (uint64, error) {
	var p [8]byte
	_, err := reader.Read(p[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(p[:]), nil
}

func DecodeByteArray(reader io.Reader, p []byte) error {
	_, err := reader.Read(p[:])
	return err
}

func DecodeSlice(reader io.Reader) ([]byte, error) {
	length, err := DecodeU32(reader)
	if err != nil {
		return nil, err
	}
	p := make([]byte, length)
	_, err = reader.Read(p)
	return p, err
}

func DecodeString(reader io.Reader) (string, error) {
	slice, err := DecodeSlice(reader)
	if err != nil {
		return "", err
	}
	return string(slice), nil
}
