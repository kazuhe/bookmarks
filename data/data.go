package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"time"

	// PostgreSQLのデータベースドライバ
	_ "github.com/lib/pq"
)

// User ユーザーを表す構造体
type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	TwitterID string    `json:"twitter_id"`
	IsPublic  bool      `json:"is_public"`
}

// DB データベースへのハンドルであり、データベース接続のプールを表す
var DB *sql.DB

// init 初期化関数でデータベースのハンドルを生成
func init() {
	var err error

	// 'sql.Open'は単にその後のDBへの接続に必要になる構造体を設定するだけでデータベースに接続する訳ではない
	DB, err = sql.Open("postgres", "user=kazuhe dbname=bookmarks sslmode=disable")
	if err != nil {
		log.Fatalf("Error openig database: %q", err)
	}
}

// createUUID "RFC4122"に基づくUUIDを作成
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt "SHA-1"を使用して160ビットのハッシュ値を生成
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

// Create 新規ユーザーの登録
func (user *User) Create() (err error) {
	// SQLのプリペアドステートメント（レコード作成時に特定の値を当てはめることができる）の定義
	statement := "insert into users (user_id, name, email, password, created_at, twitter_id, is_public) values ($1, $2, $3, $4, $5, $6, $7) returning user_id, created_at"
	// ステートメントをプリペアドステートメントとして作成するためにDB.Prepareに渡す
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// プリペアドステートメントを実行
	// 'QueryRow'で構造体sql.Row（最初の1つだけの）を返す, 'Scan'は行の中の値を引数にコピーする
	// つまり、'user.Name'等をDBに保存した後にSQLクエリによって返されたフィールドの値を構造体Userにスキャンしている
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now(), user.TwitterID, user.IsPublic).Scan(&user.UserID, &user.CreatedAt)
	return
}

// Retrive nameを元にDBから1件のUserを取得
func Retrive(name string) (user User, err error) {
	user = User{}
	// SQLのselectコマンドを使って取得したデータ（id, content, author）をpostに参照渡し
	err = DB.QueryRow("select user_id, name, email, password, created_at, twitter_id, is_public from users where name = $1", name).Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.TwitterID, &user.IsPublic)
	return
}
