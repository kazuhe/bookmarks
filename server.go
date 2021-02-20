package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/kazuhe/bookmarks/data"
)

func main() {
	server := http.Server{
		Addr: ":" + "8080",
	}

	// /post/へのリクエストをハンドラ関数'handleRequest'へリダイレクト
	// 全てのハンドラ関数は第1引数に'ResponseWriter'をとり、
	// 第2引数に'Request'をとるので改めて引数を渡す必要はない
	http.HandleFunc("/post/", handleRequest)
	log.Println("start http listenig :8080")
	server.ListenAndServe()

	// tmp:ユーザー登録を実行
	user := data.User{
		Name:     "kazuhe",
		Email:    "test@example.com",
		Password: "pass123",
	}
	if err := user.Create(); err != nil {
		log.Println("Cannot created user", err)
	}
}

// handleRequest リクエストを正しい関数に振り分けるためのハンドラ
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 検証のためにリクエストに含まれる情報を出力
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
}
