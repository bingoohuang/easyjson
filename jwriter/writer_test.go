package jwriter

import (
	stdbase64 "encoding/base64"
	"github.com/cristalhq/base64"
	"math/rand"
	"testing"
	"time"
)

var token = func() []byte {
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, 1024)
	rand.Read(token)
	return token
}()

func BenchmarkWriter_Base64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdbase64.StdEncoding.EncodeToString(token)
	}
}

func BenchmarkWriter_Cristalhq(b *testing.B) {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(token)))

	for i := 0; i < b.N; i++ {
		base64.StdEncoding.Encode(buf, token)
	}
}

/*
🕙[2021-11-12 22:21:40.046] ❯ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/easyjson/jwriter
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkWriter_Base64-12                  78285             15907 ns/op           16447 B/op          0 allocs/op
BenchmarkWriter_Cristalhq-12              431198              2782 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/bingoohuang/easyjson/jwriter 2.730s
*/
