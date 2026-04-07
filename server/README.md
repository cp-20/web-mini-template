# server

Go + Echo + MySQL のバックエンドです。

## できること

- `POST /initialize` でスキーマ再作成と初期データ投入
- feed / members / tasks API を提供
- `X-Forwarded-User` middleware を SOFT/HARD で切り替え可能
- testcontainers を使った e2e テストを実行可能

## 開発起動

```bash
mise run server
```

`server` タスクは Docker Compose の `server` サービスを起動します。
`mysql` は依存として起動されます。

## Build / Test

```bash
mise run server-build
mise run server-test
```

`server-test` では `testcontainers-go` による MySQL コンテナ e2e を含みます。

## Migration

```bash
mise run migrate-up
mise run migrate-down
```

## 環境変数

- `API_ADDR` (default: `:8080`)
- `DB_DSN` (default: `app:app@tcp(localhost:3306)/app?parseTime=true&multiStatements=true`)
- `AUTH_MODE` (`SOFT` or `HARD`, default: `SOFT`)
- `ASSETS_DIR` (指定時は静的配信)

## 主なディレクトリ

- `cmd/api`: エントリポイント
- `internal/handler`: API handler 実装
- `internal/middleware/authx`: 認証関連 middleware
- `migrations`: DB migration
- `internal/gen/openapi`: OpenAPI 生成物
