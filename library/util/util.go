package util

import (
	"context"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	GSHA1 = "gsha1"
	MD5   = "md5"
)

func Encrypt(str string) string {
	var encryptStr string
	types, _ := g.Config("config.toml").Get(context.Background(), "user.encrypt")
	switch types.String() {
	case GSHA1:
		encryptStr = gsha1.Encrypt(str)
	case MD5:
		encryptStr, _ = gmd5.Encrypt(str)
	}
	return encryptStr
}
