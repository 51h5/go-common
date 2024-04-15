package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

func WritePid(p string) error {
	if len(p) == 0 {
		return errors.New("pid file path invalid")
	}

	var err error

	d := path.Dir(p)
	err = os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		return err
	}

	return nil
}

func CopyBody(r *http.Request) ([]byte, error) {
	var dup io.ReadCloser
	r.Body, dup = DupReadCloser(r.Body)
	body, err := io.ReadAll(r.Body)
	r.Body = dup
	return body, err
}

func DupReadCloser(reader io.ReadCloser) (io.ReadCloser, io.ReadCloser) {
	var buf bytes.Buffer
	tee := io.TeeReader(reader, &buf)
	return io.NopCloser(tee), io.NopCloser(&buf)
}
