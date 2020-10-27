package utils

import "testing"

func TestIsFileExists(t *testing.T) {
    ok, err := IsFileExists("/")

    if ok {
        return
    }

    if !ok || err != nil {
        t.Fatalf("Test Exists fail, %s", err)
    }
}