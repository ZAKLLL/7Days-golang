package main

//
//import (
//	"fmt"
//	"geecache"
//	"log"
//	"net/http"
//	"testing"
//)
//
//var db = map[string]string{
//	"Tom":  "630",
//	"Jack": "589",
//	"Sam":  "567",
//}
//
//func createGroup() *geecache.Group {
//	return geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
//		func(key string) ([]byte, error) {
//			log.Println("[SlowDB] search key", key)
//			if v, ok := db[key]; ok {
//				return []byte(v), nil
//			}
//			return nil, fmt.Errorf("%s not exist", key)
//		}))
//}
//
//func startCacheServer(addr string, addrs []string, gee *geecache.Group) {
//	peers := geecache.NewHTTPPool(addr)
//	peers.Set(addrs...)
//	gee.RegisterPeers(peers)
//	log.Println("geecache is running at", addr)
//	log.Fatal(http.ListenAndServe(addr[7:], peers))
//}
//
//func startAPIServer(apiAddr string, gee *geecache.Group) {
//	http.Handle("/api", http.HandlerFunc(
//		func(w http.ResponseWriter, r *http.Request) {
//			key := r.URL.Query().Get("key")
//			view, err := gee.Get(key)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//			w.Header().Set("Content-Type", "application/octet-stream")
//			w.Write(view.ByteSlice())
//
//		}))
//	log.Println("fontend server is running at", apiAddr)
//	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
//
//}
//
//func Test_server(t *testing.T) {
//
//	apiAddr := "http://localhost:9999"
//	addrMap := map[int]string{
//		8001: "http://localhost:8001",
//		8002: "http://localhost:8002",
//		8003: "http://localhost:8003",
//	}
//
//	var addrs []string
//	for _, v := range addrMap {
//		addrs = append(addrs, v)
//	}
//	group := createGroup()
//	go startAPIServer(apiAddr, group)
//
//	go startCacheServer(addrMap[8002], addrs, createGroup())
//	go startCacheServer(addrMap[8001], addrs, createGroup())
//	startCacheServer(addrMap[8003], addrs, group)
//}
