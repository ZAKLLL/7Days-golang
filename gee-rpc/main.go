package main

import (
	"gee-rpc/server"
	"log"
	"net"
	"reflect"
	"strings"
	"sync"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	server.Accept(l)
}

//func main() {
//	addr := make(chan string)
//	go startServer(addr)
//
//	// in fact, following code is like a simple geerpc client
//	conn, _ := net.Dial("tcp", <-addr)
//	defer func() { _ = conn.Close() }()
//
//	time.Sleep(time.Second)
//	// 告知服务端本次链接的options
//	//_ = json.NewEncoder(conn).Encode(server.DefaultOption)
//	_ = gob.NewEncoder(conn).Encode(server.DefaultOption)
//
//	//使用GobCodec 作为编辑编解码器
//	cc := codec.NewGobCodec(conn)
//	// send request & receive response
//	for i := 0; i < 5; i++ {
//		h := &codec.Header{
//			ServiceMethod: "Foo.Sum",
//			Seq:           uint64(i),
//		}
//		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
//		_ = cc.ReadHeader(h)
//		var reply string
//		_ = cc.ReadBody(&reply)
//		log.Println("reply:", reply)
//	}
//}

//func main() {
//	log.SetFlags(0)
//	addr := make(chan string)
//	go startServer(addr)
//	client, _ := geerpcClient.Dial("tcp", <-addr)
//	defer func() { _ = client.Close() }()
//
//	time.Sleep(time.Second)
//	// send request & receive response
//	var wg sync.WaitGroup
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			args := fmt.Sprintf("geerpc req %d", i)
//			var reply string
//			if err := client.Call("Foo.Sum", args, &reply); err != nil {
//				log.Fatal("call Foo.Sum error:", err)
//			}
//			log.Println("reply:", reply)
//		}(i)
//	}
//	wg.Wait()
//}

func main() {
	var wg sync.WaitGroup
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		// j 从 1 开始，第 0 个入参是 wg 自己。
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}
