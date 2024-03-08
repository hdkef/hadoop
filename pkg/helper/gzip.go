package helper

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func Compress(data []byte) ([]byte, error) {
	var compressedData bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressedData)
	_, err := gzipWriter.Write(data)
	if err != nil {
		return nil, err
	}
	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}
	return compressedData.Bytes(), nil
}

func Decompress(data []byte) ([]byte, error) {
	gzipReader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	decompressedData, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}
	if err := gzipReader.Close(); err != nil {
		return nil, err
	}
	return decompressedData, nil
}
