package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"path"

	"github.com/kazuhe/bookmarks/models"
)

// UsersHandler リクエストを正しい関数に振り分けるためのハンドラ
func UsersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("UsersHandler()")

	// 検証のためにリクエストに含まれる情報を出力
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))

	var err error

	// リクエストメソッドに応じてそれぞれのCRUD関数に作業を振り分ける
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
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

// handleGet GETリクエストに応じてユーザーを返す関数
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// URLのパスを抽出
	// id, err := strconv.Atoi(path.Base(r.URL.Path))
	name := path.Base(r.URL.Path)
	// if err != nil {
	// 	return
	// }

	// メソッドRetriveでnameを元にDBの値を取得して構造体Userを作成
	user, err := models.Retrive(name)
	if err != nil {
		return
	}

	// 構造体UserをJSONフォーマットのバイト列に変換
	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	// バイト列をResponseWriterに書き出す
	w.Header().Set("Content-Type", "application/json")

	// TODO 一時的に全てのオリジンからのアクセスを許可
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(output)
	return
}

// handlePost POSTリクエストに応じてユーザーを作成する関数
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	// コンテンツの長さをサイズとしたバイト列を作成
	len := r.ContentLength
	body := make([]byte, len)

	// コンテンツ(JSON)を読み込む
	r.Body.Read(body)

	// コンテンツ(JSON)を構造体Userに組み換える
	var user models.User
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
