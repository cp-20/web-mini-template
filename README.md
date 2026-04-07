# web-mini-template

春ハッカソン向けの Web テンプレートです。

- Monorepo: `client` (Vue) + `server` (Go)
- Runtime 管理: `mise`
- Frontend: Vue 3 + Vue Router + Native CSS
- Backend: Go + Echo + sqlx + MySQL
- OpenAPI: `openapi/openapi.yaml` を source of truth として client/server を codegen
- Frontend API client: `openapi-fetch` + `openapi-typescript`
- Frontend quality: `oxfmt` + `oxlint`
- 認証 middleware: `X-Forwarded-User` を SOFT/HARD モードで処理
- Frontend/Backend ともに watch 対応

## Directory

- `client`: フロントエンド
- `server`: バックエンド
- `openapi`: OpenAPI 定義
- `.github/workflows/ci.yml`: CI
- `Dockerfile`: 統一 Dockerfile (dev/prod 両対応)
- `docker-compose.yml`: 開発用 compose

## Setup

```bash
mise trust
mise install
pnpm --dir client install
```

ルート `package.json` / `pnpm-workspace.yaml` は使っていません。

## OpenAPI Codegen

```bash
mise run codegen
```

実行される内容:

- client: `openapi-typescript` で `client/src/gen/api-types.ts` を生成
- client runtime: `openapi-fetch` が生成型を使って API 呼び出し
- server: `ogen` で `server/internal/gen/openapi` を生成

## Frontend Format/Lint

```bash
mise run client-fmt
mise run client-fmt-check
mise run client-lint
```

- `client-fmt*`: `oxfmt`
- `client-lint`: `oxlint`

## Development

### 推奨: Docker Compose

```bash
mise run dev
```

起動サービス:

- mysql: `localhost:3306`
- server: `localhost:8080` (air watch)

frontend はローカルで起動します:

```bash
mise run client
```

必要な場合のみ full docker 開発を使えます:

```bash
mise run dev-full
```

停止:

```bash
mise run dev-down
```

### 個別起動

```bash
mise run client
mise run server
mise run server-build
mise run server-test
```

`mise run server` は Docker Compose の server サービスを起動します。

## Auth Middleware

`X-Forwarded-User` を処理する middleware を提供しています。

- `SOFT` (default): ヘッダがあれば context に保存、なければ通す
- `HARD`: ヘッダがなければ `401` を返す

環境変数:

- `AUTH_MODE=SOFT|HARD`

## ASSETS_DIR 配信

`ASSETS_DIR` が設定されている場合、サーバーがそのディレクトリを静的配信します。

- 例: `ASSETS_DIR=/app/assets`
- 単一コンテナデプロイ時は `Dockerfile` の `app` stage で `client/dist` を `/app/assets` に配置済み

## Docker

### 開発 (backend + DB)

```bash
docker compose up --build mysql server
```

### 開発 (fullstack Docker)

```bash
docker compose --profile fullstack up --build
```

### 単一コンテナ (デプロイ想定)

```bash
docker build -t web-mini-template:app .
docker run --rm -p 8080:8080 \
  -e DB_DSN='app:app@tcp(host.docker.internal:3306)/app?parseTime=true&multiStatements=true' \
  -e AUTH_MODE=SOFT \
  web-mini-template:app
```

## API

- `POST /initialize`
- `GET /api/feed`
- `GET /api/members`
- `POST /api/tasks`

## CI

GitHub Actions で client/server の以下を job 分割して実行します。

- client-codegen-check:
  - client/server の codegen を実行
  - `git diff --exit-code` で生成物差分チェック
- client-fmt:
  - Format check (`oxfmt --check`)
- client-lint:
  - Lint (`oxlint --deny-warnings`)
- client-typecheck:
  - TypeCheck (`vue-tsc --noEmit`)
- client-build:
  - Build (`vp build`)
- server-build:
  - `go build ./...`
- server-test:
  - `go test ./... -v`
  - `testcontainers-go` で MySQL コンテナを起動する e2e シナリオを実行

`pnpm/action-setup` と `actions/setup-node` の cache を使って高速化しています。
