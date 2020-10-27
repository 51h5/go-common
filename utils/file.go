package utils

import (
    "errors"
    "io/ioutil"
    "os"
)

func GetFileContent(path string) ([]byte, error) {
    if path == "" {
        return nil, errors.New("path invalid")
    }

    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    return ioutil.ReadAll(f)
}

func IsFileExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return true, err
}
