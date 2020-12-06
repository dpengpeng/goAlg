package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type JsonRes struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	TimeStamp int64       `json:"timestamp"`
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	path, method := r.URL.Path, r.Method
	if path == "/" {
		w.Write([]byte("index"))
		return
	}

	if path == "/hello" && method == "POST" {
		sayHello(w, r)
		return
	}

	if path == "/sleep" {
		//mock timeout
		time.Sleep(4 * time.Second)
		return
	}
	http.Error(w, "you lost???", http.StatusNotFound)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	apiResult(w, 0, name+"say"+r.PostFormValue("some"), "success")
}

func apiResult(w http.ResponseWriter, code int, data interface{}, msg string) {
	body, _ := json.Marshal(JsonRes{Code: code, Data: data, Msg: msg, TimeStamp: time.Now().Unix()})
	w.Write(body)
}

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: http.TimeoutHandler(http.HandlerFunc(defaultHandler), 2*time.Second, "timeout!!!"),
	}
	srv.ListenAndServe()
}
