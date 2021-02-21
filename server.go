package main

import (
	"encoding/json"
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

	// /users/へのリクエストをハンドラ関数'handleRequest'へリダイレクト
	// 全てのハンドラ関数は第1引数に'ResponseWriter'をとり、
	// 第2引数に'Request'をとるので改めて引数を渡す必要はない
	http.HandleFunc("/users/", handleRequest)
	log.Println("start http listenig :8080")
	server.ListenAndServe()
}

// handleRequest リクエストを正しい関数に振り分けるためのハンドラ
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 検証のためにリクエストに含まれる情報を出力
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))

	var err error

	// リクエストメソッドに応じてそれぞれのCRUD関数に作業を振り分ける
	switch r.Method {
	// case "GET":
	// 	err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
		// case "PUT":
		// 	err = handlePut(w, r)
		// case "DELETE":
		// 	err = handleDelete(w, r)
	}

	// リクエスト自体に関わるエラー処理
	// エラーがあれば詳細とステータス500を返す
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handlePost POSTリクエストに応じて投稿を作成する関数
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	// コンテンツの長さをサイズとしたバイト列を作成
	len := r.ContentLength
	body := make([]byte, len)

	// コンテンツ(JSON)を読み込む
	r.Body.Read(body)

	// コンテンツ(JSON)を構造体Userに組み換える
	var user data.User
	json.Unmarshal(body, &user)

	// メソッドCreateで構造体UserをDBに保存
	err = user.Create()
	if err != nil {
		log.Println("Cannot created user", err)
		return
	}

	// ステータス200を返す
	w.WriteHeader(200)
	return
}
