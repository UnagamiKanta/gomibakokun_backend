# ごみばこくん バックエンド設計書

## 使用技術
プログラミング言語: Go\
フレームワーク:echo\
DB:GCP Firestore


## ディレクトリ構成
```
rootdir
├── main.go
│       
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
│       └── user_response.go
├── infrastructure
│   └── persistence
│       └── user.go
└── usecase
    └── user.go
```

## 設計(レイヤード+DDD)

### Domain layer
--- 
**ドメインモデル**
- TrashCan
  - ID(string)
  - Lattitude(Float64)精度はこのくらいで大丈夫か？実験する
  - Longtitude(Float64)
  - Image(string)変わる可能性あり
  - ImageUrl(string)
  - TrashType([]string)
    - burnable
    - unburnable
    - plastic
    - bottle_can
    - other
    - everything

## Repository layer

### Trashcan repository
domain/repository/trashcan_repository.go
trashcanに関する機能を定義(中身は実装しない)
| 関数 | 引数(型) |返り値(型)|機能(簡単に)|
| ---- | ---- | ---- | ---- |
| CreateTrashcan | ctx, trashcan | error | ゴミ箱情報をDBに保存|
| GetAllTrashcan | ctx | []trashcan, error | DBに保存してあるすべてのゴミ箱を保存 |

## Infrastructure layer
実際のDBとの接続を実装
### Trashcan persistence
infrastructure/persistence.go
| 関数 | 引数(型) |返り値(型)|機能(簡単に)|
| ---- | ---- | ---- | ---- |
| CreateTrashcan | ctx, trashcan | error | ゴミ箱情報をDBに保存|
| GetAllTrashcan | ctx | []trashcan, error | DBに保存してあるすべてのゴミ箱を保存 |

## Usecase layer
ビジネスロジックを実装
| 関数 | 引数(型) |返り値(型)|機能(簡単に)|
| ---- | ---- | ---- | ---- |
| CreateTrashcan | ctx, latitude(float64), longitude(float64), image(string), trashType([]string), nearestBuilding(string) | error | ゴミ箱のIDを作成してDBに保存|
| GetAllTrashcan | ctx | []trashcan, error | DBに保存してあるすべてのゴミ箱を保存 |