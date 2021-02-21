# bookmarks
Bookmarks app

## API Document(:tmp)

### ▼ Create User
ユーザーを新規で登録する。

ENDPOINTS:
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

ENDPOINTS:
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
