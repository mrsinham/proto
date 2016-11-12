package parser

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	type fields struct {
		r *bufio.Reader
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "simple scan",
			fields: fields{
				r: bufio.NewReader(strings.NewReader(`panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xb code=0x1 addr=0x28 pc=0x438298]

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

goroutine 1 [chan receive]:
testing.RunTests(0x5c9c28, 0x705aa0, 0x1, 0x1, 0x1)
	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:472 +0x8d5
testing.Main(0x5c9c28, 0x705aa0, 0x1, 0x1, 0x70dee0, ...)
	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:403 +0x84
main.main()
	github.com/dgryski/go-shardedkv/storage/redis/_test/_testmain.go:47 +0x9c

goroutine 4 [syscall]:
runtime.goexit()
	/home/dgryski/work/src/cvs/go/src/pkg/runtime/proc.c:1394`)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				r: tt.fields.r,
			}
			for {
				gotTok, gotLit := s.Scan()
				fmt.Println(gotTok.String(), gotLit)
				if gotTok == EOF {
					break
				}
			}

		})
	}
}
