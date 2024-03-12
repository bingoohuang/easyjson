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

func TestS2(t *testing.T) {
	r := S2{A: 150}
	data, _ := easyjson.Marshal(r)
	t.Log((string)(data))

	d := `{"a":"150","b":"true"}`
	var s2 S2
	easyjson.Unmarshal([]byte(d), &s2)
	t.Log(s2)

	d3 := `{"a":150}`
	var s3 S2
	easyjson.Unmarshal([]byte(d3), &s3)
	t.Log(s3)
}
