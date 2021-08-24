package zip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func GZipDecompress(input []byte) ([]byte, error) {
	buf := bytes.NewBuffer(input)
	reader, gzipErr := gzip.NewReader(buf)
	if gzipErr != nil {
		return nil, gzipErr
	}
	defer reader.Close()
	result, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return nil, readErr
	}
	return result, nil
}

func GZipCompress(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write(input)
	if err != nil {
		return nil, err
	}

	err = gz.Flush()
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
