package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func ChangeFileName(filename string) string {
	var f string
	m := md5.New()
	f1 := strings.Split(filename, string('.'))
	f = f1[0]
	if len(f1) > 2 {
		for i := 1; i < len(f1) - 1; i++ {
			f = f + "." + f1[i]
		}
	}
	suffix := f1[len(f1)-1]

	m.Write([]byte(f))
	fn := hex.EncodeToString(m.Sum(nil)) + "." + suffix
	return fn
}
