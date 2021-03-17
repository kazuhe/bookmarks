# bookmarks
Bookmarks app

## TODO
- user登録時に既に存在するnameがPOSTされた場合はエラーを返却
- 

## Development usage
Building a local version

```bash
$ go build -o bookmarks cmd/bookmarks/main.go
```

## API Document(:tmp)
APIドキュメント（整備中）

### ▼ Create User
下記リストの情報を元にユーザーを新規で登録する。
- name: ユーザー名
- email: メールアドレス
- password: パスワード
- twitter_id: ツイッターID
- is_public: ブックマークの公開/非公開設定

ENDPOINT:
```bash
POST /v1/users
```

SAMPLE REQUEST:
```bash
$ curl \
  -i \
  -X POST \
  -H "Content-Type: application/json" \
  http://0.0.0.0:8080/v1/users/ \
  -d '{"name": "kazuhe", "email": "kazuhe@example.com", "password": "pass1234", "twitter_id": "@kazuhe__", "is_public": "1"}'
```

### ▼ Get User
Nameを使ってユーザーを取得する。

ENDPOINT:
```bash
GET /v1/users/:id
```

SAMPLE REQUEST:
```bash
$ curl -i -X GET http://0.0.0.0:8080/v1/users/kazuhe
```

SAMPLE RESPONSE:
```json
{
	"user_id": "fdf5588f-ed66-4d1a-4dfc-8383f90cb3c0",
	"name": "kazuhe",
	"email": "kazuhe@example.com",
	"password": "789b49606c321c8cf228d17942608eff0ccc4171",
	"created_at": "2021-02-27T23:07:56.296096Z",
	"twitter_id": "@kazuhe__",
	"is_public": false
}
```

## Database design
### 要件
- ユーザーは自分専用のブックマークをn個数作成できる
- ユーザーは自分専用のタグをn個数作成できる
- ブックマークにタグをn個数設定できる
- ユーザーはブックマークの公開or非公開を選択できる
- ユーザー登録していないとブックマークもタグも作成することはできない

### テーブル
「※」印を持つテーブルを「主キー」とし、サンプルとしてテーブルに仮データを設定している。

__▼ users__
| ※user_id | name | email | password | created_at | twitter_id | is_public |
| --- | --- | --- | --- | --- | --- | --- |
| kazuhe | かずひ | kazuhe@example.com | 4a27b3ae456b0a3f7ae14e8d0b0847549b711859 | 2021-02-21 10:06:16.128659 | @kazuhe__ | true |
| betty | Betty | betty@example.com | 789b49606c321c8cf228d17942608eff0ccc4171 | 2021-02-23 12:06:20.9751 | @betty0123 | false |

__▼ bookmark__
| ※user_id | ※bookmark_id | url | comment | read_later |
| --- | --- | --- | --- | --- |
| kazuhe | 56e6c4c3c2f1 | http://example1.com | コメント | true |
| kazuhe | 5f757f0e05ae | http://example2.com | コメント | false |
| betty | fc809ffd0af0 | http://example3.com | コメント | true |
| betty | 3f9ebf8d29sg | http://example4.com | NULL | true |

__▼ tag__
| ※user_id | ※tag_id | tag |
| --- | --- | --- |
| kazuhe | develop | 開発 |
| kazuhe | life | 生活 |
| betty | life | 生活 |

__▼ bookmark_tag__
| ※user_id | ※bookmark_id | ※tag_id |
| --- | --- | --- |
| kazuhe | 56e6c4c3c2f1 | develop |
| kazuhe | 5f757f0e05ae | life |
| kazuhe | 5f757f0e05ae | develop |
| betty | fc809ffd0af0 | life |
| betty | 3f9ebf8d29sg | life |

### ER図
<img src="https://user-images.githubusercontent.com/57878514/109179839-a7136800-77cd-11eb-812f-56ef2fd8ee9d.png" width="500" alt="bookmarks ER図">
