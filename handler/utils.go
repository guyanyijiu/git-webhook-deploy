package handler

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"os"
	"os/exec"
	"syscall"
)

func HmacSha1(b []byte, secret []byte) string {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func PathIsWritable(path string) bool {
	err := syscall.Access(path, syscall.O_RDWR)
	if err != nil {
		return false
	}
	return true
}

func ExecCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
