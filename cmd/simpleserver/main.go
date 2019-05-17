package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Resp struct {
	Time        time.Time
	Random      string
	Name        string
	Age         int32
	Description string
}

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	resp := &Resp{
		Age:         12,
		Time:        time.Now(),
		Random:      RandomString(12),
		Name:        "todd",
		Description: "this is services/poi/search",
	}

	respBytes, _ := json.Marshal(resp)

	fmt.Fprintf(w, string(respBytes)) //这个写入到w的是输出到客户端的
	// fmt.Fprintf(w, string("wtf"))
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	resp := &Resp{
		Age:         12,
		Time:        time.Now(),
		Random:      RandomString(12),
		Name:        "todd",
		Description: "this is root handle",
	}
	respBytes, _ := json.Marshal(resp)

	fmt.Fprintf(w, string(respBytes)) //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/services/poi/search", sayhelloName) //设置访问的路由
	http.HandleFunc("/", rootHandle)                      //设置访问的路由
	err := http.ListenAndServe(":7171", nil)              //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
