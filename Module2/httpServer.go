package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func receiveReqHeader(w http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			fmt.Printf("%s = %s\n", k, v)
			w.Header().Set(k, v[0])
		}
	}
}

func getEnv(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "go1.19.3")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
}

func getIp(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("error:", err)
	}
	if net.ParseIP(host) != nil {
		fmt.Println("ip:", host)
	}
	fmt.Println("http code:", http.StatusOK)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200"))
}

func main() {

	// 1 接收客户端 request，并将 request 中带的 header 写入 response header
	http.HandleFunc("/work1", receiveReqHeader)

	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	http.HandleFunc("/work2", getEnv)

	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	http.HandleFunc("/work3", getIp)

	//当访问 localhost/healthz 时，应返回 200
	http.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
