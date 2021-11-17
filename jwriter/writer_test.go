package jwriter

import (
	"github.com/cristalhq/base64"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var token = func() []byte {
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, 1024*150)
	rand.Read(token)
	return token
}()

func BenchmarkWriter_Base64(b *testing.B) {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(token))+2)

	for i := 0; i < b.N; i++ {
		w := Writer{}
		w.Base64Bytes(token)
		w.Buffer.BuildBytes(buf)
	}
}

func BenchmarkWriter_Base64Old(b *testing.B) {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(token))+2)

	for i := 0; i < b.N; i++ {
		w := Writer{}
		w.Base64BytesOld(token)
		w.Buffer.BuildBytes(buf)
	}
}

func BenchmarkWriter_Cristalhq(b *testing.B) {
	p := sync.Pool{
		New: func() interface{} { return make([]byte, base64.StdEncoding.EncodedLen(len(token))) },
	}

	for i := 0; i < b.N; i++ {
		buf := p.Get().([]byte)
		base64.StdEncoding.Encode(buf, token)
		p.Put(buf)
	}
}

/*
🕙[2021-11-18 01:41:41.017] ❯ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/easyjson/jwriter
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkWriter_Base64Fast-12              27987             42330 ns/op             151 B/op          2 allocs/op
BenchmarkWriter_Base64-12                   3782            282393 ns/op         1103995 B/op         25 allocs/op
BenchmarkWriter_Cristalhq-12               27216             44016 ns/op              31 B/op          1 allocs/op
PASS
ok      github.com/bingoohuang/easyjson/jwriter 4.372s
*/
