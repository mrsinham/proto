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
//		{
//			args: args{
//				r: strings.NewReader(`
//goroutine 3 [running]:
//runtime.panic(0x550400, 0x70ad88)
//	/home/dgryski/work/src/cvs/go/src/pkg/runtime/panic.c:266 +0xb6
//testing.func·005()
//	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:385 +0xe8
//runtime.panic(0x550400, 0x70ad88)
//	/home/dgryski/work/src/cvs/go/src/pkg/runtime/panic.c:248 +0x106
//github.com/dgryski/go-shardedkv/storage/redis.(*Storage).Get(0xc21004d120, 0x57efa0, 0x5, 0x7febbe4eada8, 0x1, ...)
//	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storage/redis/redis.go:29 +0xd8
//github.com/dgryski/go-shardedkv/storagetest.StorageTest(0xc21004e090, 0x7febbe65f580, 0xc21004d120)
//	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storagetest/storagetest.go:42 +0x550
//github.com/dgryski/go-shardedkv/storage/redis.TestRedis(0xc21004e090)
//	/home/dgryski/Dropbox/GITS/gocode/src/github.com/dgryski/go-shardedkv/storage/redis/redis_test.go:16 +0x135
//testing.tRunner(0xc21004e090, 0x705aa0)
//	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:391 +0x8b
//created by testing.RunTests
//	/home/dgryski/work/src/cvs/go/src/pkg/testing/testing.go:471 +0x8b2
//				`),
//			},
//		},
		{
			args: args{
				r: strings.NewReader(`
panic: open /var/folders/z0/vp96rcnd63j6_xlslhplwm4r0000gp/T/proxy-service/55982bb21a0f599d45005412: too many open files

goroutine 15238 [running]:
runtime.panic(0x2a8340, 0xc21011cf90)
        /usr/local/go/src/pkg/runtime/panic.c:266 +0xb6
main.NewFileStream(0xc210ae2c00, 0x57, 0xc21004e2c0, 0x339030, 0xc, ...)
        /Users/giannimoschini/src/github.com/leibowitz/go-proxy-service/proxy.go:461 +0x76
main.func·002(0xc21011db40, 0xc210f525b0, 0xce0dc)
        /Users/giannimoschini/src/github.com/leibowitz/go-proxy-service/proxy.go:363 +0x2f1
github.com/leibowitz/goproxy.FuncRespHandler.Handle(0xc2100881e0, 0xc21011db40, 0xc210f525b0, 0xc21011db40)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/actions.go:35 +0x36
github.com/leibowitz/goproxy.func·016(0xc21011db40, 0xc210f525b0, 0xcaef5)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/dispatcher.go:279 +0x16d
github.com/leibowitz/goproxy.FuncRespHandler.Handle(0xc210088200, 0xc21011db40, 0xc210f525b0, 0xc21011db40)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/actions.go:35 +0x36
github.com/leibowitz/goproxy.(*ProxyHttpServer).filterResponse(0xc21004caf0, 0xc21011db40, 0xc210f525b0, 0xc21011db40)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:72 +0x7b
github.com/leibowitz/goproxy.func·018()
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:207 +0x9f0
created by github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:254 +0x1075

goroutine 1 [sleep]:
time.Sleep(0x4c4b40)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/time.goc:31 +0x31
net/http.(*Server).Serve(0xc210085fa0, 0x742dc8, 0xc210000db8, 0x0, 0x0)
        /usr/local/go/src/pkg/net/http/server.go:1634 +0x1f9
net/http.(*Server).ListenAndServe(0xc210085fa0, 0xc210085fa0, 0x742da0)
        /usr/local/go/src/pkg/net/http/server.go:1612 +0xa0
net/http.ListenAndServe(0x7fff5fbfeb84, 0x5, 0x742da0, 0xc21004caf0, 0x1, ...)
        /usr/local/go/src/pkg/net/http/server.go:1677 +0x6d
main.main()
        /Users/giannimoschini/src/github.com/leibowitz/go-proxy-service/proxy.go:420 +0xb92

goroutine 3 [select]:
labix.org/v2/mgo.(*mongoCluster).syncServersLoop(0xc21008d000)
        /Users/giannimoschini/src/labix.org/v2/mgo/cluster.go:366 +0x52c
created by labix.org/v2/mgo.newCluster
        /Users/giannimoschini/src/labix.org/v2/mgo/cluster.go:72 +0x120

goroutine 15252 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc210d69690, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210d69660, 0x29, 0xc2101761c0, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc2120a07e0, 0xc210d69630, 0x21, 0x79bac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc2120a07e0, 0x35b370, 0x1b, 0x79bac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210a3eaa0, 0xc210dccb60)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210a3eaa0, 0xc210dccb60)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210a3eaa0, 0xc210dccb60)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc2102e8c80)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 5 [syscall]:
runtime.goexit()
        /usr/local/go/src/pkg/runtime/proc.c:1394

goroutine 7 [runnable]:
net.runtime_pollWait(0x742ba0, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc21004c840, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc21004c840, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc21004c7e0, 0xc210089f60, 0x24, 0x24, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc210000b80, 0xc210089f60, 0x24, 0x24, 0x0, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
labix.org/v2/mgo.fill(0x741bd8, 0xc210000b80, 0xc210089f60, 0x24, 0x24, ...)
        /Users/giannimoschini/src/labix.org/v2/mgo/socket.go:489 +0x5b
labix.org/v2/mgo.(*mongoSocket).readLoop(0xc2100498c0)
        /Users/giannimoschini/src/labix.org/v2/mgo/socket.go:506 +0x115
created by labix.org/v2/mgo.newSocket
        /Users/giannimoschini/src/labix.org/v2/mgo/socket.go:163 +0x2b3

goroutine 8 [sleep]:
time.Sleep(0x12a05f200)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/time.goc:31 +0x31
labix.org/v2/mgo.(*mongoServer).pinger(0xc210049700, 0xc210049701)
        /Users/giannimoschini/src/labix.org/v2/mgo/server.go:284 +0x10f
created by labix.org/v2/mgo.newServer
        /Users/giannimoschini/src/labix.org/v2/mgo/server.go:87 +0xf6

goroutine 418 [select]:
net/http.(*persistConn).writeLoop(0xc2100f1c80)
        /usr/local/go/src/pkg/net/http/transport.go:791 +0x271
created by net/http.(*Transport).dialConn
        /usr/local/go/src/pkg/net/http/transport.go:529 +0x61e
goroutine 3843 [IO wait]:
net.runtime_pollWait(0xf274f8, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc2106bf0d0, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc2106bf0d0, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc2106bf070, 0xc210a36000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc2105c4700, 0xc210a36000, 0x1000, 0x1000, 0x247d00, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
net/http.(*liveSwitchReader).Read(0xc21020d3a8, 0xc210a36000, 0x1000, 0x1000, 0xc210474720, ...)
        /usr/local/go/src/pkg/net/http/server.go:204 +0xa5
io.(*LimitedReader).Read(0xc2101805a0, 0xc210a36000, 0x1000, 0x1000, 0x12f17, ...)
        /usr/local/go/src/pkg/io/io.go:398 +0xbb
bufio.(*Reader).fill(0xc21037cea0)
        /usr/local/go/src/pkg/bufio/bufio.go:91 +0x110
bufio.(*Reader).ReadSlice(0xc21037cea0, 0x1370a, 0x0, 0x0, 0x0, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:274 +0x204
bufio.(*Reader).ReadLine(0xc21037cea0, 0x0, 0x0, 0x0, 0x0, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:305 +0x63
net/textproto.(*Reader).readLineSlice(0xc21011c450, 0x738000, 0x2d3020, 0xc21011c450, 0x27d02, ...)
        /usr/local/go/src/pkg/net/textproto/reader.go:55 +0x61
net/textproto.(*Reader).ReadLine(0xc21011c450, 0xc2102c40d0, 0x0, 0xc210a37000, 0xc2106bf7e0)
        /usr/local/go/src/pkg/net/textproto/reader.go:36 +0x27
net/http.ReadRequest(0xc21037cea0, 0xc2102c40d0, 0x0, 0x0)
        /usr/local/go/src/pkg/net/http/request.go:526 +0x88
net/http.(*conn).readRequest(0xc21020d380, 0x0, 0x0, 0x0)
        /usr/local/go/src/pkg/net/http/server.go:575 +0x1bb
net/http.(*conn).serve(0xc21020d380)
        /usr/local/go/src/pkg/net/http/server.go:1123 +0x3b4
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 15281 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc210ef0840, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210ef07b0, 0x29, 0xc21070d940, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc2107e8b60, 0xc210ef0780, 0x21, 0x791ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc2107e8b60, 0x35b370, 0x1b, 0x791ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210854b40, 0xc210af5d00)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210854b40, 0xc210af5d00)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210854b40, 0xc210af5d00)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210a46280)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b
goroutine 417 [IO wait]:
net.runtime_pollWait(0x742708, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc21017cc30, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc21017cc30, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc21017cbd0, 0xc2103de800, 0x400, 0x400, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc2103c4658, 0xc2103de800, 0x400, 0x400, 0x0, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
crypto/tls.(*block).readFromUntil(0xc2100c3360, 0x742e00, 0xc2103c4658, 0x5, 0xc2103c4658, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:459 +0xb6
crypto/tls.(*Conn).readRecord(0xc210452000, 0x17, 0x0, 0x8)
        /usr/local/go/src/pkg/crypto/tls/conn.go:539 +0x107
crypto/tls.(*Conn).Read(0xc210452000, 0xc210488000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:897 +0x135
bufio.(*Reader).fill(0xc21013f1e0)
        /usr/local/go/src/pkg/bufio/bufio.go:91 +0x110
bufio.(*Reader).Peek(0xc21013f1e0, 0x1, 0x0, 0x0, 0x0, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:119 +0xcb
net/http.(*persistConn).readLoop(0xc2100f1c80)
        /usr/local/go/src/pkg/net/http/transport.go:687 +0xb7
created by net/http.(*Transport).dialConn
        /usr/local/go/src/pkg/net/http/transport.go:528 +0x607
goroutine 15257 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc2109b88d0, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc2109b8630, 0x29, 0xc2104db6c0, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc210f52850, 0xc2109b8540, 0x21, 0x78dac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc210f52850, 0x35b370, 0x1b, 0x78dac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210dcd1e0, 0xc21055eea0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210dcd1e0, 0xc21055eea0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210dcd1e0, 0xc21055eea0)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc2102b9800)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 15296 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc21073c210, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc21073c1e0, 0x29, 0xc2108f4240, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc2107e8ee0, 0xc21073c1b0, 0x21, 0x78bac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc2107e8ee0, 0x35b370, 0x1b, 0x78bac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210806c80, 0xc210b890d0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210806c80, 0xc210b890d0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210806c80, 0xc210b890d0)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210a46a80)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 15303 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc21073cea0, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc21073ce70, 0x29, 0xc21070d6e0, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc211dae070, 0xc21073ce40, 0x21, 0x789ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc211dae070, 0x35b370, 0x1b, 0x789ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc2106310a0, 0xc2104139c0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc2106310a0, 0xc2104139c0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc2106310a0, 0xc2104139c0)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210a46e00)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 69 [IO wait]:
net.runtime_pollWait(0x7425b8, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc2101fcf40, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc2101fcf40, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc2101fcee0, 0xc2101d0f30, 0x24, 0x24, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc21012be60, 0xc2101d0f30, 0x24, 0x24, 0x0, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
labix.org/v2/mgo.fill(0x741bd8, 0xc21012be60, 0xc2101d0f30, 0x24, 0x24, ...)
        /Users/giannimoschini/src/labix.org/v2/mgo/socket.go:489 +0x5b
labix.org/v2/mgo.(*mongoSocket).readLoop(0xc2102070e0)
        /Users/giannimoschini/src/labix.org/v2/mgo/socket.go:506 +0x115
created by labix.org/v2/mgo.newSocket
        /Users/giannimoschini/src/labix.org/v2/mgo/socket.go:163 +0x2b3

goroutine 15297 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc21073c3c0, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc21073c390, 0x29, 0xc2108f43a0, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc210710620, 0xc21073c360, 0x21, 0x785ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc210710620, 0x35b370, 0x1b, 0x785ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210806d20, 0xc210b894e0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210806d20, 0xc210b894e0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210806d20, 0xc210b894e0)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210a46b00)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 468 [IO wait]:
net.runtime_pollWait(0x742900, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc2106130d0, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc2106130d0, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc210613070, 0xc210614000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc2103c42c0, 0xc210614000, 0x1000, 0x1000, 0x247d00, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
net/http.(*liveSwitchReader).Read(0xc2101673a8, 0xc210614000, 0x1000, 0x1000, 0xc21025b2a0, ...)
        /usr/local/go/src/pkg/net/http/server.go:204 +0xa5
io.(*LimitedReader).Read(0xc210484540, 0xc210614000, 0x1000, 0x1000, 0x12f17, ...)
        /usr/local/go/src/pkg/io/io.go:398 +0xbb
bufio.(*Reader).fill(0xc2104a9240)
        /usr/local/go/src/pkg/bufio/bufio.go:91 +0x110
bufio.(*Reader).ReadSlice(0xc2104a9240, 0x340a, 0x0, 0x0, 0x0, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:274 +0x204
bufio.(*Reader).ReadLine(0xc2104a9240, 0x0, 0x0, 0x0, 0x0, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:305 +0x63
net/textproto.(*Reader).readLineSlice(0xc2102ec4e0, 0x738000, 0x23a1a0, 0xb61ce8, 0x27d02, ...)
        /usr/local/go/src/pkg/net/textproto/reader.go:55 +0x61
net/textproto.(*Reader).ReadLine(0xc2102ec4e0, 0xc21045d8f0, 0x0, 0xc210615000, 0xc210613230)
        /usr/local/go/src/pkg/net/textproto/reader.go:36 +0x27
net/http.ReadRequest(0xc2104a9240, 0xc21045d8f0, 0x0, 0x0)
        /usr/local/go/src/pkg/net/http/request.go:526 +0x88
net/http.(*conn).readRequest(0xc210167380, 0x0, 0x0, 0x0)
        /usr/local/go/src/pkg/net/http/server.go:575 +0x1bb
net/http.(*conn).serve(0xc210167380)
        /usr/local/go/src/pkg/net/http/server.go:1123 +0x3b4
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 3529 [IO wait]:
net.runtime_pollWait(0x762c60, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc21034f680, 0x72, 0x741148, 0x23)

        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc21034f680, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc21034f620, 0xc21057e000, 0x400, 0x400, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc2104c8da8, 0xc21057e000, 0x400, 0x400, 0x1b, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
crypto/tls.(*block).readFromUntil(0xc210347f90, 0x742e00, 0xc2104c8da8, 0x5, 0xc2104c8da8, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:459 +0xb6
crypto/tls.(*Conn).readRecord(0xc210151000, 0x17, 0x0, 0x1b2ed)
        /usr/local/go/src/pkg/crypto/tls/conn.go:539 +0x107
crypto/tls.(*Conn).Read(0xc210151000, 0xc210747000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:897 +0x135
bufio.(*Reader).fill(0xc2100e7480)
        /usr/local/go/src/pkg/bufio/bufio.go:91 +0x110
bufio.(*Reader).Peek(0xc2100e7480, 0x1, 0x2, 0x2, 0x2, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:119 +0xcb
github.com/leibowitz/goproxy.isEof(0xc2100e7480, 0xc210151000)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:49 +0x30
github.com/leibowitz/goproxy.func·018()
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:176 +0x27b
created by github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:254 +0x1075

goroutine 15301 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc21073cb40, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc21073cb10, 0x29, 0xc21070d180, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc210a71e70, 0xc21073cae0, 0x21, 0xb7dac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc210a71e70, 0x35b370, 0x1b, 0xb7dac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210854be0, 0xc21087da90)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210854be0, 0xc21087da90)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210854be0, 0xc21087da90)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210a46d00)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 15295 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc21073c060, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc21073c030, 0x29, 0xc2108f4060, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc2107e8770, 0xc210ef0bd0, 0x21, 0xb7bac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc2107e8770, 0x35b370, 0x1b, 0xb7bac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210806be0, 0xc210b89dd0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210806be0, 0xc210b89dd0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210806be0, 0xc210b89dd0)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210a46a00)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 15284 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc210a11210, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210a111e0, 0x29, 0xc210a70340, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc210710770, 0xc210a111b0, 0x21, 0xb79ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc210710770, 0x35b370, 0x1b, 0xb79ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210806000, 0xc210b89820)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210806000, 0xc210b89820)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210806000, 0xc210b89820)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210a46380)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 15251 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc210d69480, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210d69420, 0x29, 0xc210181c20, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc2120a0380, 0xc210d69390, 0x21, 0xba1ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc2120a0380, 0x35b370, 0x1b, 0xba1ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210a3ea00, 0xc210dcca90)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210a3ea00, 0xc210dcca90)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210a3ea00, 0xc210dcca90)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc2102e8b80)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 15323 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc210732cf0, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210732cc0, 0x29, 0xc21070d760, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc211dae380, 0xc210732c90, 0x21, 0x7c1ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc211dae380, 0x35b370, 0x1b, 0x7c1ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210631280, 0xc21045d0d0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210631280, 0xc21045d0d0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210631280, 0xc21045d0d0)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210c7c580)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 120 [IO wait]:
net.runtime_pollWait(0x742318, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc21021bd10, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc21021bd10, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc21021bcb0, 0xc21026a800, 0x400, 0x400, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc21012b070, 0xc21026a800, 0x400, 0x400, 0x1b, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
crypto/tls.(*block).readFromUntil(0xc2102da2d0, 0x742e00, 0xc21012b070, 0x5, 0xc21012b070, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:459 +0xb6
crypto/tls.(*Conn).readRecord(0xc210151a00, 0x17, 0x0, 0x1b2ed)
        /usr/local/go/src/pkg/crypto/tls/conn.go:539 +0x107
crypto/tls.(*Conn).Read(0xc210151a00, 0xc2101b3000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:897 +0x135
bufio.(*Reader).fill(0xc2100f8e40)
        /usr/local/go/src/pkg/bufio/bufio.go:91 +0x110
bufio.(*Reader).Peek(0xc2100f8e40, 0x1, 0x2, 0x2, 0x2, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:119 +0xcb
github.com/leibowitz/goproxy.isEof(0xc2100f8e40, 0xc210151a00)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:49 +0x30
github.com/leibowitz/goproxy.func·018()
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:176 +0x27b
created by github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:254 +0x1075

goroutine 15254 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc210bc60c0, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210d69fc0, 0x29, 0xc210592400, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc2120a0e00, 0xc210d69f90, 0x21, 0xb75ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc2120a0e00, 0x35b370, 0x1b, 0xb75ac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210a3ec80, 0xc210dccdd0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210a3ec80, 0xc210dccdd0)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210a3ec80, 0xc210dccdd0)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc2102e8d00)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 3073 [IO wait]:
net.runtime_pollWait(0x742a50, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc2107015a0, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc2107015a0, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc210701540, 0xc210753800, 0x400, 0x400, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc2104c8fd0, 0xc210753800, 0x400, 0x400, 0x1b, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
crypto/tls.(*block).readFromUntil(0xc2103476c0, 0x742e00, 0xc2104c8fd0, 0x5, 0xc2104c8fd0, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:459 +0xb6
crypto/tls.(*Conn).readRecord(0xc2107e1780, 0x17, 0x0, 0x1b2ed)
        /usr/local/go/src/pkg/crypto/tls/conn.go:539 +0x107
crypto/tls.(*Conn).Read(0xc2107e1780, 0xc2103b6000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:897 +0x135
bufio.(*Reader).fill(0xc2103eb9c0)
        /usr/local/go/src/pkg/bufio/bufio.go:91 +0x110
bufio.(*Reader).Peek(0xc2103eb9c0, 0x1, 0x2, 0x2, 0x2, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:119 +0xcb
github.com/leibowitz/goproxy.isEof(0xc2103eb9c0, 0xc2107e1780)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:49 +0x30
github.com/leibowitz/goproxy.func·018()
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:176 +0x27b
created by github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:254 +0x1075

goroutine 15243 [runnable]:
math/big.nat.mul(0xc211e91280, 0x8, 0x14, 0xc211e911e0, 0x8, ...)
        /usr/local/go/src/pkg/math/big/nat.go:379
math/big.nat.expNNWindowed(0xc211e911e0, 0x8, 0x14, 0xc2101fab40, 0x8, ...)
        /usr/local/go/src/pkg/math/big/nat.go:1354 +0xa16
math/big.nat.expNN(0xc2101faba0, 0x8, 0xc, 0xc2101fab40, 0x8, ...)
        /usr/local/go/src/pkg/math/big/nat.go:1255 +0x3b1
math/big.nat.probablyPrime(0xc210234000, 0x8, 0xc, 0x14, 0xc)
        /usr/local/go/src/pkg/math/big/nat.go:1440 +0x848
math/big.(*Int).ProbablyPrime(0xc211ebec00, 0x14, 0xc21004a920)
        /usr/local/go/src/pkg/math/big/int.go:721 +0x4a
crypto/rand.Prime(0x743288, 0xc211e0f5f0, 0x200, 0xc211ebec00, 0x0, ...)
        /usr/local/go/src/pkg/crypto/rand/util.go:97 +0x353
crypto/rsa.GenerateMultiPrimeKey(0x743288, 0xc211e0f5f0, 0x2, 0x400, 0xc210234f60, ...)
        /usr/local/go/src/pkg/crypto/rsa/rsa.go:166 +0x20e
crypto/rsa.GenerateKey(0x743288, 0xc211e0f5f0, 0x400, 0x14, 0x0, ...)
        /usr/local/go/src/pkg/crypto/rsa/rsa.go:125 +0x56
github.com/leibowitz/goproxy.signHost(0xc2100701a0, 0x1, 0x1, 0x2bf2a0, 0xc2100394e0, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/signer.go:74 +0x822
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210a3e460, 0xc210625a90)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:155 +0xccf
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210a3e460, 0xc210625a90)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210a3e460, 0xc210625a90)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc2102e8500)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 3085 [IO wait]:
net.runtime_pollWait(0x741c88, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc2102e51b0, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc2102e51b0, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc2102e5150, 0xc2106be000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0
net.(*conn).Read(0xc210343210, 0xc2106be000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/net/net.go:122 +0xc5
crypto/tls.(*block).readFromUntil(0xc210347c30, 0x742e00, 0xc210343210, 0x5, 0xc210343210, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:459 +0xb6
crypto/tls.(*Conn).readRecord(0xc21027b780, 0x17, 0x0, 0x0)
        /usr/local/go/src/pkg/crypto/tls/conn.go:539 +0x107
crypto/tls.(*Conn).Read(0xc21027b780, 0xc2106bdb25, 0x4db, 0x4db, 0x0, ...)
        /usr/local/go/src/pkg/crypto/tls/conn.go:897 +0x135
bufio.(*Reader).fill(0xc21016cf00)
        /usr/local/go/src/pkg/bufio/bufio.go:91 +0x110
bufio.(*Reader).ReadSlice(0xc21016cf00, 0x17e0000000a, 0x0, 0x0, 0x0, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:274 +0x204
bufio.(*Reader).ReadLine(0xc21016cf00, 0x0, 0x0, 0x0, 0x0, ...)
        /usr/local/go/src/pkg/bufio/bufio.go:305 +0x63
net/textproto.(*Reader).readLineSlice(0xc2102622d0, 0x738000, 0x23a1a0, 0xb77c70, 0x27d02, ...)
        /usr/local/go/src/pkg/net/textproto/reader.go:55 +0x61
net/textproto.(*Reader).ReadLine(0xc2102622d0, 0xc21087dea0, 0x0, 0x0, 0x0)
        /usr/local/go/src/pkg/net/textproto/reader.go:36 +0x27
net/http.ReadRequest(0xc21016cf00, 0xc21087dea0, 0x0, 0x0)
        /usr/local/go/src/pkg/net/http/request.go:526 +0x88
github.com/leibowitz/goproxy.func·018()
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:177 +0x298
created by github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:254 +0x1075

goroutine 15292 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc210d9a1b0, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210d9a180, 0x29, 0xc2104dc360, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc210d9ca10, 0xc210d9a150, 0x21, 0xbbdac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc210d9ca10, 0x35b370, 0x1b, 0xbbdac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210806960, 0xc210b89410)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210806960, 0xc210b89410)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210806960, 0xc210b89410)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc210a46800)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 15264 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc210af4f90, 0x27, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210af4f60, 0x29, 0xc21053bfc0, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc213a1fd90, 0xc210af4e70, 0x21, 0xbbbac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc213a1fd90, 0x35b370, 0x1b, 0xbbbac8, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps(0xc21004caf0, 0x742ef8, 0xc210dcd6e0, 0xc21055e340)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:80 +0x35f
github.com/leibowitz/goproxy.(*ProxyHttpServer).ServeHTTP(0xc21004caf0, 0x742ef8, 0xc210dcd6e0, 0xc21055e340)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/proxy.go:99 +0xa1
net/http.serverHandler.ServeHTTP(0xc210085fa0, 0x742ef8, 0xc210dcd6e0, 0xc21055e340)
        /usr/local/go/src/pkg/net/http/server.go:1597 +0x16e
net/http.(*conn).serve(0xc2102b9600)
        /usr/local/go/src/pkg/net/http/server.go:1167 +0x7b7
created by net/http.(*Server).Serve
        /usr/local/go/src/pkg/net/http/server.go:1644 +0x28b

goroutine 7199 [semacquire]:
sync.runtime_Semacquire(0xc210085f54)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/sema.goc:199 +0x30
sync.(*Mutex).Lock(0xc210085f50)
        /usr/local/go/src/pkg/sync/mutex.go:66 +0xd6
log.(*Logger).Output(0xc210085f50, 0x2, 0xc21089d1c0, 0x1b8, 0x0, ...)
        /usr/local/go/src/pkg/log/log.go:134 +0x95
log.(*Logger).Printf(0xc210085f50, 0xc210181180, 0x15, 0xc2101810e0, 0x2, ...)
        /usr/local/go/src/pkg/log/log.go:160 +0x7a
github.com/leibowitz/goproxy.(*ProxyCtx).printf(0xc210373a80, 0xc2106200f0, 0xd, 0xbb9e10, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:47 +0x1fc
github.com/leibowitz/goproxy.(*ProxyCtx).Logf(0xc210373a80, 0x324040, 0x7, 0xbb9e10, 0x1, ...)
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/ctx.go:60 +0xa4
github.com/leibowitz/goproxy.func·018()
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:192 +0x615
created by github.com/leibowitz/goproxy.(*ProxyHttpServer).handleHttps
        /Users/giannimoschini/src/github.com/leibowitz/goproxy/https.go:254 +0x1075

goroutine 3847 [IO wait]:
net.runtime_pollWait(0x7634e8, 0x72, 0x0)
        /private/tmp/makerelease863497612/go/src/pkg/runtime/netpoll.goc:116 +0x6a
net.(*pollDesc).Wait(0xc21028dd10, 0x72, 0x741148, 0x23)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:81 +0x34
net.(*pollDesc).WaitRead(0xc21028dd10, 0x23, 0x741148)
        /usr/local/go/src/pkg/net/fd_poll_runtime.go:86 +0x30
net.(*netFD).Read(0xc21028dcb0, 0xc2105be000, 0x1000, 0x1000, 0x0, ...)
        /usr/local/go/src/pkg/net/fd_unix.go:204 +0x2a0

				`),
			},
		},
	}
	for _, tt := range tests {
		p := NewParser(tt.args.r)

		p.Parse()
	}
}
