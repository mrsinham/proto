package parser

import (
	"io"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				r: strings.NewReader(`
goroutine 3 [running]:
runtime.panic(0x550400, 0x70ad88)
	/home/dgryski/work/src/cvs/go/src/pkg/runtime/panic.c:266 +0xb6
testing.func·005()
	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:385 +0xe8
runtime.panic(0x550400, 0x70ad88)
	/home/dgryski/work/src/cvs/go/src/pkg/runtime/panic.c:248 +0x106
github.com/dgryski/go-shardedkv/storage/redis.(*Storage).Get(0xc21004d120, 0x57efa0, 0x5, 0x7febbe4eada8, 0x1, ...)
	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storage/redis/redis.go:29 +0xd8
github.com/dgryski/go-shardedkv/storagetest.StorageTest(0xc21004e090, 0x7febbe65f580, 0xc21004d120)
	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storagetest/storagetest.go:42 +0x550
github.com/dgryski/go-shardedkv/storage/redis.TestRedis(0xc21004e090)
	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storage/redis/redis_test.go:16 +0x135
testing.tRunner(0xc21004e090, 0x705aa0)
	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:391 +0x8b
created by testing.RunTests
	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:471 +0x8b2
				`),
			},
		},
	}
	for _, tt := range tests {
		p := NewParser(tt.args.r)

		var tok Token
		tok, _ = p.scan()
		if tok != NewLine {
			t.Fatal("must have NewLine")
		}

		tok, _ = p.scan()
		if tok != Goroutine {
			t.Fatal("must have goroutine")
		}

		p.unscan()
		tok, _ = p.scan()
		if tok != Goroutine {
			t.Fatal("must have been goroutine")
		}
		p.unscan()
		p.unscan()
		tok, _ = p.scan()

		if tok != NewLine {
			t.Fatal("must have been NewLine")
		}
	}
}

func TestParser_Parse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				r: strings.NewReader(`
goroutine 3 [running]:
runtime.panic(0x550400, 0x70ad88)
	/home/dgryski/work/src/cvs/go/src/pkg/runtime/panic.c:266 +0xb6
testing.func·005()
	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:385 +0xe8
runtime.panic(0x550400, 0x70ad88)
	/home/dgryski/work/src/cvs/go/src/pkg/runtime/panic.c:248 +0x106
github.com/dgryski/go-shardedkv/storage/redis.(*Storage).Get(0xc21004d120, 0x57efa0, 0x5, 0x7febbe4eada8, 0x1, ...)
	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storage/redis/redis.go:29 +0xd8
github.com/dgryski/go-shardedkv/storagetest.StorageTest(0xc21004e090, 0x7febbe65f580, 0xc21004d120)
	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storagetest/storagetest.go:42 +0x550
github.com/dgryski/go-shardedkv/storage/redis.TestRedis(0xc21004e090)
	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storage/redis/redis_test.go:16 +0x135
testing.tRunner(0xc21004e090, 0x705aa0)
	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:391 +0x8b
created by testing.RunTests
	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:471 +0x8b2
				`),
			},
		},
	}
	for _, tt := range tests {
		p := NewParser(tt.args.r)

		p.Parse()
	}
}
