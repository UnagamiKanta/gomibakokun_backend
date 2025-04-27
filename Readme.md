# ごみばこくん バックエンド設計書

## 使用技術
プログラミング言語: Go\
フレームワーク:echo\
DB:GCP Firestore





## ディレクトリ構成
```
rootdir
├── cmd
│   └── api
│       └── main.go
├── domain
│   └── repository
|   |   └── user_repository.go
|   └── user.go
├── config //DBの起動など
│   └── database.go
├── interfaces
│   └── handler
│   |   └── user.go
|   └── response
│       └── response.go
├── infrastructure
│   └── persistence
│       └── user.go
└── usecase
    └── user.go
```

## 設計(レイヤード+DDD)

### domain layer
--- 
**ドメインモデル**
- TrashCan
  - ID(string)
  - Lattitude(Float64)精度はこのくらいで大丈夫か？実験する
  - Longtitude(Float64)
  - NearestBuilding(string)変わる可能性あり
  - ImageUrl(string)
  - TrashType([]string)
    - burnable
    - unburnable
    - plastic
    - bottle_can
    - other
    - everything
  
**リポジトリ**
型が明らかなものは省略(これってDB参照してる時点で不味くね？)
| 関数 | 引数(型) |返り値(型)|
| ---- | ---- | ----|
| Create | clinet, tarashCan |error|
| Read | DB(*firesotre.Clinet), ID(string) | trashCan, error|
|Delete|DB, ID(string)|error|
