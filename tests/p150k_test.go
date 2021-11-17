package tests

import (
	"github.com/bingoohuang/easyjson"
	"github.com/cristalhq/base64"
	"testing"
)

func Test150K(t *testing.T) {
	r := Rsp{A: make([]byte, 150*1024)}
	data, _ := easyjson.Marshal(r)
	if len(data) != base64.StdEncoding.EncodedLen(len(r.A))+8 {
		t.Fail()
	}
}
