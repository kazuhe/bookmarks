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
「※」印を持つテーブルを「キー」とする。

__▼ ユーザー__
| ※ユーザーID | メールアドレス | パスワード | 登録日 | Twitterアカウント情報 | 公開/非公開設定 |
| --- | --- | --- | --- | --- | --- |
| kazuhe | kazuhe@example.com | 4a27b3ae456b0a3f7ae14e8d0b0847549b711859 | 2021-02-21 10:06:16.128659 | @kazuhe__ | true |
| betty | betty@example.com | 789b49606c321c8cf228d17942608eff0ccc4171 | 2021-02-21 12:06:20.9751 | @kazuhe__ | false |

__▼ ブックマーク__
| ※ユーザーID | ※ブックマークID | URL | コメント | タグID | あとで読むフラグ |
| --- | --- | --- | --- | --- | --- |
| kazuhe | 1 | http://example1.com | コメント | T1 | true |
| kazuhe | 2 | http://example2.com | コメント | T1 | false |
| kazuhe | 2 | http://example2.com | | T2 | false |
| betty | 1 | http://example3.com | コメント | T1 | true |
| betty | 2 | http://example4.com | | T1 | true |

__▼ タグ__
| ※ユーザーID | ※タグID | タグ |
| --- | --- | --- |
| kazuhe | T1 | develop |
| kazuhe | T2 | life |
| betty | T1 | life |

### ER図
![Screen Shot 2021-02-24 at 23 32 48](https://user-images.githubusercontent.com/57878514/109015610-a6f56880-76f8-11eb-8416-62fff70985fb.png)