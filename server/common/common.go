package common

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
)


//Gzip编码
func GzipEncode(b []byte) ([]byte, error) {
	var buf bytes.Buffer

	w := gzip.NewWriter(&buf)
	defer w.Close()

	_, err := w.Write(b)
	if err != nil {
		return []byte{}, err
	}

	err = w.Flush()
	if err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

//Gzip解码
func GzipDecode(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	buf.Write(b)

	r, err := gzip.NewReader(&buf)
	if err != nil {
		return []byte{}, err
	}
	defer r.Close()

	ret, err := ioutil.ReadAll(r)
	if err != nil && err != io.ErrUnexpectedEOF {
		return []byte{}, err
	}

	return ret, nil
}
