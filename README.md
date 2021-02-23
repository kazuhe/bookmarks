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

### 論理設計
1. エンティティの抽出
   - 利用ユーザー
     - id
     - メールアドレス
     - パスワード
     - 登録日
     - twitterアカウント情報
     - GitHubアカウント情報（?）
     - 公開/非公開設定
   - ブックマークデータ
     - コメント
     - 独自設定可能なタグ
     - あとで読むフラグ
2. エンティティの定義
3. 正規化
4. ER図の作成

### 物理設計
1. テーブル定義
2. インデックス定義
3. ハードウェアのサイジング
4. ストレージの冗長構成決定
5. ファイルの物理配置決定

