# bookmarks
Bookmarks app

## API Document(:tmp)
APIドキュメント（整備中）

### ▼ Create User
ユーザーを新規で登録する。

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
  http://0.0.0.0:8080/users/ \
  -d '{"name": "kazuhe", "email": "kazuhe@example.com","password": "pass1234"}'
```

### ▼ Get User
IDを使ってユーザーを取得する。

ENDPOINT:
```bash
POST /v1/users/:id
```

SAMPLE REQUEST:
```bash
$ curl -i -X GET http://0.0.0.0:8080/users/1
```

```json
{
  "id": 1,
	"uuid": "0555868c-ec17-41c3-55dc-56e6c4c3c2f1",
	"name": "kazuhe",
	"email": "kazuhe@example.com",
	"password": "4a27b3ae456b0a3f7ae14e8d0b0847549b711859",
	"created_at": "2021-02-21T10:06:16.128659Z"
}
```

## Database design
DB設計（整備中）

### 正規化したテーブル
「※」印を持つテーブルを「主キー」とする。

__▼ ユーザー__
| ※ユーザーID | メールアドレス | パスワード | 登録日 | Twitterアカウント情報 | 公開/非公開設定 |
| --- | --- | --- | --- | --- | --- |
| kazuhe | kazuhe@example.com | 4a27b3ae456b0a3f7ae14e8d0b0847549b711859 | 2021-02-21 10:06:16.128659 | @kazuhe__ | true |
| betty | betty@example.com | 789b49606c321c8cf228d17942608eff0ccc4171 | 2021-02-21 12:06:20.9751 | @kazuhe__ | false |

__▼ ブックマーク__
| ※ユーザーID | ※ブックマークID | URL | コメント | あとで読むフラグ |
| --- | --- | --- | --- | --- |
| kazuhe | B1 | http://example1.com | コメント | true |
| kazuhe | B2 | http://example2.com | コメント | false |
| kazuhe | B2 | http://example2.com | | false |
| betty | B1 | http://example3.com | コメント | true |
| betty | B2 | http://example4.com | | true |

__▼ タグ__
| ※ユーザーID | ※タグID | タグ |
| --- | --- | --- |
| kazuhe | T1 | develop |
| kazuhe | T2 | life |
| betty | T1 | life |

__▼ ブックマークタグ__
| ※ユーザーID | ※タグID | ※ブックマークID |
| --- | --- | --- |
| kazuhe | T1 | B1 |
| kazuhe | T1 | B2 |
| kazuhe | T2 | B2 |
| betty | T1 | B1 |
| betty | T1 | B2 |

### ER図
![Bookmarks ER図](https://user-images.githubusercontent.com/57878514/109039495-335f5580-7710-11eb-9c32-ec196eb8eae5.png)
