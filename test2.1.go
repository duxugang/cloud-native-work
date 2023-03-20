package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
*
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
*/
func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(resp http.ResponseWriter, req *http.Request) {
	// 1
	for k, v := range req.Header {
		resp.Header().Set(k, strings.Join(v, ","))
	}
	// 2
	resp.Header().Set("VERSION", os.Getenv("VERSION"))
	// 3
	log.Printf("源IP：%v", req.Host)
	// 4
	resp.WriteHeader(200)
	io.WriteString(resp, "200")
}
