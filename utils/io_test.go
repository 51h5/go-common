package utils

import (
    "io/ioutil"
    "os"
    "strconv"
    "testing"
)

func TestWritePidWithoutPath(t *testing.T) {
    p := ""

    if err := WritePid(p); err == nil {
        t.Errorf("WritePid 无效文件路径验证失败")
        return
    } else if err.Error() != "pid file path invalid" {
        t.Errorf("WritePid 无效验证错误信息不匹配")
        return
    }
}

func TestWritePidWithPath(t *testing.T) {
    pid := strconv.Itoa(os.Getpid())
    p := "./io.pid"

    if err := WritePid(p); err != nil {
        t.Errorf("WritePid 写入失败: %s", err)
        return
    }
    defer os.Remove(p)

    if b, err := ioutil.ReadFile(p); err != nil {
        t.Errorf("pid 临时文件读取失败: %s", err)
        return
    } else if string(b) != pid {
        t.Errorf("pid 写入不匹配: %s !== %s", pid, b)
        return
    }
}
