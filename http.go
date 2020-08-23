package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var ServerPort int

func Serve() {
	//监听协议
	http.HandleFunc("/", Handler)
	//监听服务
	err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(ServerPort), nil)
	if err != nil {
		fmt.Printf("服务器错误，端口（%d）可能被占用", ServerPort)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if !(r.Method == "GET" && r.URL.Path == "/") {
		w.WriteHeader(403)
		return
	}
	keyword := r.FormValue("kw")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(ToJson(Transform(Parse(BookmarkFilePath), keyword)))
}
